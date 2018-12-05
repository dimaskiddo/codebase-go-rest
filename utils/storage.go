package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Storage Configuration Struct
type storageS3Config struct {
	DisableSSL bool
	Endpoint   string
	AccessKey  string
	SecretKey  string
	Region     string
	Bucket     string
}

// Storage Configuration Variable
var storageS3Cfg storageS3Config

// Storage Connection Variable
var StorageS3 *session.Session

// Storage Connect Function
func initStorageS3() {
	var config *aws.Config

	// Create S3 Connection
	switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		config = &aws.Config{
			Credentials: credentials.NewStaticCredentials(storageS3Cfg.AccessKey, storageS3Cfg.SecretKey, ""),
			Region:      aws.String(storageS3Cfg.Region),
		}
	case "minio":
		// Set Endpoint URL Based On SSL Support
		var endpoint string
		if storageS3Cfg.DisableSSL {
			endpoint = "http://" + storageS3Cfg.Endpoint
		} else {
			endpoint = "https://" + storageS3Cfg.Endpoint
		}

		config = &aws.Config{
			Credentials:      credentials.NewStaticCredentials(storageS3Cfg.AccessKey, storageS3Cfg.SecretKey, ""),
			Endpoint:         aws.String(endpoint),
			Region:           aws.String(storageS3Cfg.Region),
			DisableSSL:       aws.Bool(storageS3Cfg.DisableSSL),
			S3ForcePathStyle: aws.Bool(true),
		}
	}

	// Return Session
	StorageS3 = session.New(config)
}

func StorageS3GetFileLink(fileName string) string {
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		// Return Composed URL Based on Storage Driver
		switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
		case "aws":
			return "https://s3." + storageS3Cfg.Region + ".amazonaws.com/" + storageS3Cfg.Bucket + "/" + fileName
		case "minio":
			if storageS3Cfg.DisableSSL {
				return "http://" + storageS3Cfg.Endpoint + "/" + storageS3Cfg.Bucket + "/" + fileName
			} else {
				return "https://" + storageS3Cfg.Endpoint + "/" + storageS3Cfg.Bucket + "/" + fileName
			}
		}
	}

	return ""
}

func StorageS3UploadFile(fileName string) {
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		// Open File for Upload
		fileContent, err := os.Open(fileName)
		if err != nil {
			log.Println("failed to open file " + fileName)
			return
		}
		defer fileContent.Close()

		// Create Uploader and Do Upload
		uploader := s3manager.NewUploader(StorageS3)
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(storageS3Cfg.Bucket),
			Key:    aws.String(filepath.Base(fileName)),
			Body:   fileContent,
		})

		// Check for Upload Error
		if err != nil {
			log.Println("failed to upload file " + fileName + ", " + err.Error())
			return
		}

		// Upload Success
		log.Println("successfully upload file " + fileName + " to " + result.Location)
	}
}
