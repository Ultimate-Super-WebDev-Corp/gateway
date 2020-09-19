package search

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/search"
)

func (s Search) SearchByIds(ctx context.Context, msg *search.SearchByIdsRequest) (*search.Product, error) {
	logger := ctxzap.Extract(ctx)

	s3GetReq, _ := s.s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.s3Bucket),
		Key:    aws.String(msg.UUIDs[0]),
	})
	s3GetReqStr, _ := s3GetReq.Presign(15 * time.Minute) //todo err

	logger.Info(s3GetReqStr)
	return &search.Product{}, nil
}
