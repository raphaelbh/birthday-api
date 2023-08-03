package photos

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetAll() []any {

	accessKey, _ := os.LookupEnv("ACCESS_KEY")
	secretKey, _ := os.LookupEnv("SECRET_KEY")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})
	if err != nil {
		exitErrorf("Unable to list items in bucket, %v", err)
	}

	// Create S3 service client
	svc := s3.New(sess)

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String("bailedajack")})
	if err != nil {
		exitErrorf("Unable to list items in bucket, %v", err)
	}

	var items []any
	for _, item := range resp.Contents {
		items = append(items, *item.Key)
	}

	return items
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
