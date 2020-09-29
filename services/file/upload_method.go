package file

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (fu File) Upload(stream file.File_UploadServer) error {
	partition := time.Now().Format("20060102")
	bucket := makeBucketName(fu.s3Bucket, partition)

	if _, err := fu.s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}); err != nil && !isBucketExists(err) {
		return server.NewErrServer(codes.Internal, errors.WithStack(err))
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
		if err == io.EOF {
			break
		}
		if err != nil {
			return server.NewErrServer(codes.Internal, errors.WithStack(err))
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
				return server.NewErrServer(codes.InvalidArgument, errors.New("the first message must be metadata"))
			}
			if _, errWrite := pw.Write(ch.Chunk); errWrite != nil {
				return server.NewErrServer(codes.InvalidArgument, errors.WithStack(errWrite))
			}

		default:
			return server.NewErrServer(codes.InvalidArgument, errors.New("unknown type of chunk"))
		}
	}

	_ = pw.Close()

	if !asyncOk {
		fileUuid, errS3 = asyncResp.get()
	}
	if errS3 != nil {
		return server.NewErrServer(codes.Internal, errors.WithStack(errS3))
	}

	return stream.SendAndClose(&file.FileUploadResponse{UUID: fileUuid})
}
