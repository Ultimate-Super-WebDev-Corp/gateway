package file

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
)

const oneFilePart = 8 * 1024 * 1024 // 1 MB

func (fu File) GetFile(msg *file.FileUUID, stream file.File_GetFileServer) error {
	partition, err := getPartition(msg.UUID)
	if err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}

	wasMeta := false
	for chFrom, chTo := 0, oneFilePart-1; ; {
		respGetObj, err := fu.s3Client.GetObject(&s3.GetObjectInput{
			Bucket: aws.String(makeBucketName(fu.s3Bucket, partition)),
			Key:    aws.String(msg.UUID),
			Range:  aws.String(fmt.Sprintf("bytes=%d-%d", chFrom, chTo)),
		})
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if !wasMeta {
			meta, err := extractMetadataFromAWS(respGetObj.Metadata)
			if err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			if err := stream.Send(&file.Chunk{
				OneOfChunk: &file.Chunk_Meta{
					Meta: meta,
				},
			}); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
			wasMeta = true
		}

		body, err := ioutil.ReadAll(respGetObj.Body)
		if err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		bodySize := len(body)

		if err := stream.Send(&file.Chunk{
			OneOfChunk: &file.Chunk_Chunk{
				Chunk: body,
			},
		}); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		if bodySize < oneFilePart {
			break
		}

		chFrom = chTo
		chTo = chFrom + oneFilePart
	}

	return nil
}
