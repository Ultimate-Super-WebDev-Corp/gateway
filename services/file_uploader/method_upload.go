package file_uploader

import (
	"io"

	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file_uploader"
)

func (fu FileUploader) Upload(stream file_uploader.FileUploader_UploadServer) error {
	pr, pw := io.Pipe()
	defer func() {
		_ = pr.Close()
	}()

	logger := ctxzap.Extract(stream.Context())

	var wasMeta bool
	var msg *file_uploader.Chunk
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
		case *file_uploader.Chunk_Meta:
			if !wasMeta {
				wasMeta = true
				asyncResp = fu.asyncUploadToS3(pr, ch.Meta)
			} else {
				logger.Error("metadata must be sent only once")
			}

		case *file_uploader.Chunk_Chunk:
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

	return stream.SendAndClose(&file_uploader.FileUploaderResponse{UUID: fileUuid})
}
