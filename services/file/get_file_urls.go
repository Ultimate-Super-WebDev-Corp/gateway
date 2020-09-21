package file

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
)

func (fu File) GetFileURLs(_ context.Context, msg *file.FileUUIDs) (*file.FileURLs, error) {

	urls := make([]string, 0, len(msg.UUIDs))
	for _, uuid := range msg.UUIDs {
		partition, err := getPartition(uuid)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		s3GetReq, _ := fu.s3Client.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(makeBucketName(fu.s3Bucket, partition)),
			Key:    aws.String(uuid),
		})
		url, err := s3GetReq.Presign(15 * time.Minute)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		urls = append(urls, url)
	}

	return &file.FileURLs{URLs: urls}, nil
}
