package file_uploader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/Ultimate-Super-WebDev-Corp/gateway/gen/services/file_uploader"
)

type FileUploader struct {
	s3Uploader *s3manager.Uploader
	s3Bucket   string
}

type config struct {
	S3Endpoint   string `env:"S3_ENDPOINT"`
	S3AccessKey  string `env:"S3_ACCESS_KEY"`
	S3SecretKey  string `env:"S3_SECRET_KEY"`
	S3Region     string `env:"S3_REGION"`
	S3BucketName string `env:"S3_BUCKET_NAME"`
}

type Dependences struct {
	Registrar *grpc.Server
}

func NewFileUploader(dep Dependences) error {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		return errors.WithStack(err)
	}

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(cfg.S3AccessKey, cfg.S3SecretKey, ""),
		Endpoint:         aws.String(cfg.S3Endpoint),
		Region:           aws.String(cfg.S3Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return errors.WithStack(err)
	}

	fu := &FileUploader{
		s3Uploader: s3manager.NewUploader(newSession),
		s3Bucket:   cfg.S3BucketName,
	}

	file_uploader.RegisterFileUploaderServer(dep.Registrar, fu)

	return nil
}
