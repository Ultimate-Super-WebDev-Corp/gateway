package file

import (
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/pkg/errors"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
)

func makeBucketName(bucket, partition string) string {
	return strings.Join([]string{bucket, partition}, ".")
}

func makeFileUUID(uuid, partition string) string {
	return strings.Join([]string{uuid, partition}, ".")
}

func getPartition(str string) (string, error) {
	strs := strings.Split(str, ".")
	if len(strs) != 2 {
		return "", errors.Errorf("cannot get partition from %s", str)
	}
	return strs[1], nil
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

func makeAWSMetadata(meta *file.FileMetadata) map[string]*string {
	return map[string]*string{
		"Type": aws.String(strconv.Itoa(int(meta.Type))),
	}
}

func extractMetadataFromAWS(meta map[string]*string) (*file.FileMetadata, error) {
	strFileType := meta["Type"]
	if strFileType == nil {
		return nil, errors.New("file type not defined")
	}
	fileType, err := strconv.Atoi(*strFileType)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &file.FileMetadata{
		Type: file.FileType(fileType),
	}, nil
}
