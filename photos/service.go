package photos

import (
	"fmt"
	"mime/multipart"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Client *s3.S3
var once sync.Once

func getS3Client() *s3.S3 {
	once.Do(func() {

		accessKey, _ := os.LookupEnv("ACCESS_KEY")
		secretKey, _ := os.LookupEnv("SECRET_KEY")

		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("us-east-2"),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
		})
		if err != nil {
			exitErrorf("Unable create a connection with amazon aws, %v", err)
		}

		s3Client = s3.New(sess)
	})

	return s3Client
}

func Upload(files []*multipart.FileHeader) (err error) {

	svc := getS3Client()

	for _, file := range files {

		// open file
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// send file
		_, err = svc.PutObject(&s3.PutObjectInput{
			Bucket: aws.String("bailedajack"),
			Key:    aws.String(file.Filename),
			Body:   src,
		})
		if err != nil {
			return err
		}

	}

	return nil
}

func GetAll() []any {

	svc := getS3Client()
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String("bailedajack")})
	if err != nil {
		exitErrorf("Unable to list items in bucket, %v", err)
	}

	items := []any{}
	for _, item := range resp.Contents {
		items = append(items, *item.Key)
	}

	return items
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
