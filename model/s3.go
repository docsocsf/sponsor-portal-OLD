package model

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"mime"
	"path"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3 struct {
	session *session.Session
	service *s3.S3

	bucket string
	prefix string
}

func NewS3(config *aws.Config, bucket, prefix string) *S3 {
	session := session.New(config)
	service := s3.New(session)

	return &S3{
		session: session,
		service: service,

		bucket: bucket,
		prefix: prefix,
	}
}

func (s *S3) Put(key string, data io.Reader) error {
	log.Printf("Putting file '%s' into '%s'\n", key, s.bucket)
	buf, err := ioutil.ReadAll(data)
	if err != nil {
		return err
	}

	_, err = s.service.PutObject(&s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(s.prefix + key),
		Body:        bytes.NewReader(buf),
		ContentType: aws.String(mime.TypeByExtension(path.Ext(key))),
	})
	return err
}

func (s *S3) Get(key string) (*s3.GetObjectOutput, error) {
	log.Printf("Getting file '%s' from '%s'\n", key, s.bucket)
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(s.prefix + key),
	}

	result, err := s.service.GetObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeNoSuchKey:
				log.Println(s3.ErrCodeNoSuchKey, aerr.Error())
				return nil, DbError{true, aerr}
			default:
				log.Println(aerr.Error())
			}
		}
		log.Println(err.Error())
		return nil, err
	}

	return result, nil
}
