package helpers

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/nutwreck/admin-pos-service/configs"
)

func NewStorageClient() (*s3.S3, error) {
	awsConfig := &aws.Config{
		Region:           aws.String(configs.Region),
		Endpoint:         aws.String(configs.Endpoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials: credentials.NewStaticCredentials(
			configs.AccessKeyID,
			configs.SecretAccessKey,
			"", // optional security token
		),
	}

	// Buat sesi AWS
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}

	// Buat dan kembalikan klien S3
	return s3.New(sess), nil
}

func UploadFileToStorageClient(fileContent *multipart.FileHeader, fileName string, acl string) error {
	// Inisialisasi klien S3 dari helper
	s3Client, err := NewStorageClient()
	if err != nil {
		fmt.Println("Error NewStorageClient => " + err.Error())
		return err
	}

	// Membuka file dari form yang sudah diberikan oleh framework
	file, err := fileContent.Open()
	if err != nil {
		fmt.Println("Error fileContent.Open => " + err.Error())
		return err
	}
	defer file.Close()

	// Mengatur parameter untuk operasi unggah
	params := &s3.PutObjectInput{
		Bucket: aws.String(configs.BucketName),
		Body:   file,
		Key:    aws.String(fileName),
		ACL:    aws.String(acl), // Atur izin sesuai kebutuhan Anda
	}

	// Melakukan unggah file
	_, err = s3Client.PutObject(params)
	if err != nil {
		fmt.Println("Error unggah file => " + err.Error())
		return err
	}

	return nil
}

func DeleteFileFromStorageClient(fileName string) error {
	// Inisialisasi klien S3 dari helper
	s3Client, err := NewStorageClient()
	if err != nil {
		return err
	}

	// Mengatur parameter untuk operasi penghapusan
	params := &s3.DeleteObjectInput{
		Bucket: aws.String(configs.BucketName),
		Key:    aws.String(fileName),
	}

	// Melakukan penghapusan file
	_, err = s3Client.DeleteObject(params)
	if err != nil {
		return err
	}

	return nil
}
