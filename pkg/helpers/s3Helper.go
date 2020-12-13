package helpers

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Helper struct {
	service    *s3.S3
	defaultACL string
}

type ObjectDescription struct {
	Key        string
	FilePath   string
	BucketName string
	ACL        string
	MetaData   map[string]string
}

func NewS3Helper(sess *session.Session, defaultACL string) *S3Helper {
	return &S3Helper{
		service:    s3.New(sess),
		defaultACL: defaultACL,
	}
}

func (h S3Helper) UploadObject(objDescription ObjectDescription) (*s3.PutObjectOutput, error) {
	f, err := os.Open(objDescription.FilePath)
	if err != nil {
		return &s3.PutObjectOutput{}, err
	}
	defer f.Close()

	input := &s3.PutObjectInput{
		Bucket:   aws.String(objDescription.BucketName),
		Key:      aws.String(objDescription.Key),
		Body:     aws.ReadSeekCloser(f),
		Metadata: rebuildMetadata(objDescription.MetaData),
	}

	result, err := h.service.PutObject(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				return &s3.PutObjectOutput{}, aerr
			}
		}
		return &s3.PutObjectOutput{}, err
	}
	return result, nil
}

func rebuildMetadata(metadata map[string]string) map[string]*string {
	var tmp map[string]*string
	for key, value := range metadata {
		tmp[key] = aws.String(value)
	}
	return tmp
}
