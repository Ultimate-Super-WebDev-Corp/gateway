package file

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
)

func (fu File) Upload(stream file.File_UploadServer) error {
	partition := time.Now().Format("20060102")
	bucket := makeBucketName(fu.s3Bucket, partition)

	if _, err := fu.s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}); err != nil && !isBucketExists(err) {
		return status.Error(codes.Internal, err.Error())
	}

	pr, pw := io.Pipe()
	defer func() {
		_ = pr.Close()
	}()

	logger := ctxzap.Extract(stream.Context())

	var wasMeta bool
	var msg *file.Chunk
	var fileUuid string
	var err, errS3 error
	var asyncOk bool
	var asyncResp asyncUploadToS3Response

	for {
		msg, err = stream.Recv()
		if err != nil {
			break
		}
		if fileUuid, errS3, asyncOk = asyncResp.getNotBlocking(); asyncOk {
			break
		}

		switch ch := msg.OneOfChunk.(type) {
		case *file.Chunk_Meta:
			if !wasMeta {
				wasMeta = true
				asyncResp = fu.asyncUploadToS3(bucket, pr, ch.Meta)
			} else {
				logger.Error("metadata must be sent only once")
			}

		case *file.Chunk_Chunk:
			if !wasMeta {
				return status.Error(codes.InvalidArgument, "the first message must be metadata")
			}
			if _, errWrite := pw.Write(ch.Chunk); errWrite != nil {
				return status.Error(codes.InvalidArgument, errWrite.Error())
			}

		default:
			return status.Error(codes.InvalidArgument, "unknown type of chunk")
		}
	}

	if err != io.EOF && err != nil {
		return status.Error(codes.Internal, err.Error())
	} else {
		_ = pw.Close()
	}

	if !asyncOk {
		fileUuid, errS3 = asyncResp.get()
	}
	if errS3 != nil {
		return status.Error(codes.Internal, errS3.Error())
	}

	return stream.SendAndClose(&file.FileUploadResponse{UUID: fileUuid})
}

func isBucketExists(err error) bool {
	if err == nil {
		return false
	}

	errAws, ok := err.(awserr.Error)
	if !ok {
		return false
	}
	return errAws.Code() == s3.ErrCodeBucketAlreadyOwnedByYou ||
		errAws.Code() == s3.ErrCodeBucketAlreadyExists
}
