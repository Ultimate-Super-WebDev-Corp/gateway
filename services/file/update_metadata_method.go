package file

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
	"github.com/Ultimate-Super-WebDev-Corp/gateway/server"
)

func (fu File) UpdateMetadata(_ context.Context, msg *file.UpdateFileMetadata) (*empty.Empty, error) {
	if msg.Meta == nil {
		return &empty.Empty{}, nil
	}

	partition, err := getPartition(msg.UUID)
	if err != nil {
		return nil, server.NewErrServer(codes.InvalidArgument, errors.WithStack(err))
	}

	bucket := makeBucketName(fu.s3Bucket, partition)
	headObjectResp, err := fu.s3Client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(msg.UUID),
	})
	if err != nil {
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
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
		return nil, server.NewErrServer(codes.Internal, errors.WithStack(err))
	}

	return &empty.Empty{}, nil
}
