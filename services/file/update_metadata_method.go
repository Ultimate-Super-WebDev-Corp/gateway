package file

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
)

func (fu File) UpdateMetadata(_ context.Context, msg *file.UpdateFileMetadata) (*empty.Empty, error) {
	if msg.Meta == nil {
		return &empty.Empty{}, nil
	}

	partition, err := getPartition(msg.UUID)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	bucket := makeBucketName(fu.s3Bucket, partition)
	headObjectResp, err := fu.s3Client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(msg.UUID),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	meta := updateAWSMetadata(headObjectResp.Metadata, msg.Meta)
	source := bucket + "/" + msg.UUID
	_, err = fu.s3Client.CopyObject(&s3.CopyObjectInput{
		CopySource:        aws.String(source),
		Bucket:            aws.String(bucket),
		Key:               aws.String(msg.UUID),
		Metadata:          meta,
		MetadataDirective: aws.String("REPLACE"),
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &empty.Empty{}, nil
}
