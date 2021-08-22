package services

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
	"path"
)

type S3Client struct {
	client *s3.Client
	bucket string
	prefix string
}

func NewS3Session(region string, bucket string, key string, secret string) (*S3Client, error) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(key, secret, "")),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	client := s3.NewFromConfig(cfg)

	return &S3Client{
		client: client,
		bucket: bucket,
		prefix: "public",
	}, nil
}

func (s *S3Client) GetS3KeyByBookID(bookId string) (string, error) {
	key := path.Join(s.prefix, fmt.Sprintf("%s.epub", bookId))
	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := s.client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucket),
		Prefix: &key,
	})
	if err != nil {
		log.Fatal(err)
	}
	if len(output.Contents) == 0 {
		return "", nil
	}
	return aws.ToString(output.Contents[0].Key), nil
}