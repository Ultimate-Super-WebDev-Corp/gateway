package file

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

const oneFilePart = 1024 * 1024 // 1 MB

func (fu File) GetFile(msg *file.FileUUID, stream file.File_GetFileServer) error {
	partition, err := getPartition(msg.UUID)
	if err != nil {
		return server.NewErrServer(codes.InvalidArgument, errors.WithStack(err))
	}

	wasMeta := false
	for chFrom, chTo := 0, oneFilePart-1; ; {
		respGetObj, err := fu.s3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(makeBucketName(fu.s3Bucket, partition)),
			Key:    aws.String(msg.UUID),
			Range:  aws.String(fmt.Sprintf("bytes=%d-%d", chFrom, chTo)),
		})
		if err != nil {
			return server.NewErrServer(codes.Internal, errors.WithStack(err))
		}

		if !wasMeta {
			meta, err := extractMetadataFromAWS(respGetObj.Metadata)
			if err != nil {
				return server.NewErrServer(codes.Internal, errors.WithStack(err))
			}

			if err := stream.Send(&file.Chunk{
				OneOfChunk: &file.Chunk_Meta{
					Meta: meta,
				},
			}); err != nil {
				return server.NewErrServer(codes.Internal, errors.WithStack(err))
			}
			wasMeta = true
		}

		body, err := ioutil.ReadAll(respGetObj.Body)
		if err != nil {
			return server.NewErrServer(codes.Internal, errors.WithStack(err))
		}

		bodySize := len(body)

		if err := stream.Send(&file.Chunk{
			OneOfChunk: &file.Chunk_Chunk{
				Chunk: body,
			},
		}); err != nil {
			return server.NewErrServer(codes.Internal, errors.WithStack(err))
		}

		if bodySize < oneFilePart {
			break
		}

		chFrom = chTo
		chTo = chFrom + oneFilePart
	}

	return nil
}
