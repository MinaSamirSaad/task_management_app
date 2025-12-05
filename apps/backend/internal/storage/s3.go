package storage

import (
	"context"
	"io"

	appconfig "github.com/MinaSamirSaad/go-tasker/internal/config"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Storage struct {
	Client *s3.Client
	Bucket string
}

func NewS3Storage(cfg *appconfig.AWSConfig) (*S3Storage, error) {
	awsCfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(cfg.Region),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				cfg.AccessKeyID,
				cfg.SecretAccessKey,
				"",
			),
		),
	)
	if err != nil {
		return nil, err
	}

	return &S3Storage{
		Client: s3.NewFromConfig(awsCfg),
		Bucket: cfg.UploadBucket,
	}, nil
}

func (s *S3Storage) UploadFile(key string, body io.Reader) error {
	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: &s.Bucket,
		Key:    &key,
		Body:   body,
	})
	return err
}
