package file

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file"
)

type asyncUploadToS3Response struct {
	err  chan error
	uuid chan string
}

func (r asyncUploadToS3Response) get() (string, error) {
	select {
	case id := <-r.uuid:
		return id, nil
	case err := <-r.err:
		return "", err
	}
}

func (r asyncUploadToS3Response) getNotBlocking() (string, error, bool) {
	select {
	case id := <-r.uuid:
		return id, nil, true
	case err := <-r.err:
		return "", err, true
	default:
		return "", nil, false
	}
}

func (fu File) asyncUploadToS3(bucket string, body io.Reader, meta *file.FileMetadata) asyncUploadToS3Response {
	resp := asyncUploadToS3Response{
		err:  make(chan error, 1),
		uuid: make(chan string, 1),
	}
	go func(bucket string, body io.Reader, resp asyncUploadToS3Response, meta *file.FileMetadata) {
		defer func() {
			if p := recover(); p != nil {
				resp.err <- errors.Errorf("recovering from panic %v", p)
			}
			close(resp.err)
			close(resp.uuid)
		}()
		if fileUuid, err := fu.uploadToS3(bucket, body, meta); err != nil {
			resp.err <- err
		} else {
			resp.uuid <- fileUuid
		}
	}(bucket, body, resp, meta)

	return resp
}

var fileTypeToContentType = map[file.FileType]string{
	file.FileType_JPEG: "image/jpeg",
}

func (fu File) uploadToS3(bucket string, body io.Reader, meta *file.FileMetadata) (string, error) {
	partition, err := getPartition(bucket)
	if err != nil {
		return "", err
	}

	contentType := fileTypeToContentType[meta.Type]
	if contentType == "" {
		return "", errors.New("unknown file type")
	}

	UUID, err := uuid.NewUUID()
	if err != nil {
		return "", errors.WithStack(err)
	}

	fileUUID := makeFileUUID(UUID.String(), partition)
	_, err = fu.s3Uploader.Upload(&s3manager.UploadInput{
		Body:        body,
		Bucket:      aws.String(bucket),
		Key:         aws.String(makeFileUUID(UUID.String(), partition)),
		Expires:     aws.Time(time.Now().UTC().Add(time.Hour)),
		ContentType: aws.String(contentType),
	})

	return fileUUID, errors.WithStack(err)
}
