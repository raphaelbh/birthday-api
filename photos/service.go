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
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var once sync.Once

var s3Client *s3.S3
var s3Uploader *s3manager.Uploader

func getS3Uploader() *s3manager.Uploader {
	once.Do(func() {

		accessKey, _ := os.LookupEnv("ACCESS_KEY")
		secretKey, _ := os.LookupEnv("SECRET_KEY")

		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("us-east-2"),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
			DisableSSL:  aws.Bool(true),
		})
		if err != nil {
			exitErrorf("Unable create a connection with amazon aws, %v", err)
		}

		s3Uploader = s3manager.NewUploader(sess)
	})

	return s3Uploader
}

func getS3Client() *s3.S3 {
	once.Do(func() {

		accessKey, _ := os.LookupEnv("ACCESS_KEY")
		secretKey, _ := os.LookupEnv("SECRET_KEY")

		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String("us-east-2"),
			Credentials: credentials.NewStaticCredentials(accessKey, secretKey, ""),
			DisableSSL:  aws.Bool(true),
		})
		if err != nil {
			exitErrorf("Unable create a connection with amazon aws, %v", err)
		}

		s3Client = s3.New(sess)
	})

	return s3Client
}

func Upload(files []*multipart.FileHeader) (err error) {

	fmt.Println("[service.Upload] Iniciando upload dos arquivos")

	uploader := getS3Uploader()
	fmt.Println("[service.Upload] svc recuperado: ", uploader)

	fmt.Println("[service.Upload] Quantidade de arquivos recebidos: ", len(files))
	fmt.Println("[service.Upload] Arquivos recebidos: ", files)
	for _, file := range files {

		// open file
		fmt.Println("[service.Upload] Abrindo arquivo (pegando source): ", file)
		src, err := file.Open()
		if err != nil {
			fmt.Println("[service.Upload] Erro ao abrir arquivo: ", err)
			return err
		}
		defer src.Close()

		// send file
		fmt.Println("[service.Upload] Enviando arquivo para bucket: ", aws.String(file.Filename))
		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String("bailedajack"),
			Key:    aws.String(file.Filename),
			Body:   src,
		})
		if err != nil {
			fmt.Println("[service.Upload] Erro ao enviar arquivo para bucket: ", err)
			return err
		}

	}

	return nil
}

func GetAll() ([]any, error) {

	fmt.Println("[photos.GetAll] Iniciando processamento")
	svc := getS3Client()
	fmt.Println("[photos.GetAll] svc recuperado: ", svc)

	fmt.Println("[photos.GetAll] Iniciando listagem")
	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String("bailedajack"),
	})
	if err != nil {
		fmt.Println("[photos.GetAll] Erro ao listar objetos do bucket: ", err)
		return nil, err
	}

	items := []any{}
	for _, item := range resp.Contents {
		items = append(items, *item.Key)
	}

	return items, nil
}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
