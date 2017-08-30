package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/caarlos0/env"
)

type s3Env struct {
	Bucket string `env:"S3_BUCKET"`

	Region   string `env:"AWS_REGION"`
	Endpoint string `env:"S3_ENDPOINT"`
}

type S3 struct {
	Aws    *aws.Config
	Bucket string
	Prefix string
}

func GetS3(prefix string) (S3, error) {
	var s3Env s3Env
	err := env.Parse(&s3Env)
	if err != nil {
		return S3{}, err
	}

	awsConfig := aws.Config{
		Endpoint:         aws.String(s3Env.Endpoint),
		Region:           aws.String(s3Env.Region),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	}

	return S3{
		Aws:    &awsConfig,
		Bucket: s3Env.Bucket,
		Prefix: prefix,
	}, nil
}
