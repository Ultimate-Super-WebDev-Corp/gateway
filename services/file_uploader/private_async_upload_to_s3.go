package file_uploader

import (
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file_uploader"
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

func (fu FileUploader) asyncUploadToS3(body io.Reader, meta *file_uploader.Metadata) asyncUploadToS3Response {
	resp := asyncUploadToS3Response{
		err:  make(chan error, 1),
		uuid: make(chan string, 1),
	}
	go func(body io.Reader, resp asyncUploadToS3Response, meta *file_uploader.Metadata) {
		defer func() {
			if p := recover(); p != nil {
				resp.err <- errors.Errorf("recovering from panic %v", p)
			}
			close(resp.err)
			close(resp.uuid)
		}()
		if fileUuid, err := fu.uploadToS3(body, meta); err != nil {
			resp.err <- err
		} else {
			resp.uuid <- fileUuid
		}
	}(body, resp, meta)

	return resp
}

var fileTypeToContentType = map[file_uploader.FileType]string{
	file_uploader.FileType_JPEG: "image/jpeg",
}

func (fu FileUploader) uploadToS3(body io.Reader, meta *file_uploader.Metadata) (string, error) {
	contentType := fileTypeToContentType[meta.Type]
	if contentType == "" {
		return "", errors.New("unknown file type")
	}

	fileUUID, err := uuid.NewUUID()
	if err != nil {
		return "", errors.WithStack(err)
	}

	_, err = fu.s3Uploader.Upload(&s3manager.UploadInput{
		Body:        body,
		Bucket:      aws.String(fu.s3Bucket),
		Key:         aws.String(fileUUID.String()),
		Expires:     aws.Time(time.Now().UTC().Add(time.Hour)),
		ContentType: aws.String(contentType),
	})

	return fileUUID.String(), errors.WithStack(err)
}
