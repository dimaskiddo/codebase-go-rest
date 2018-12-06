package utils

import (
	"errors"
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
type storeS3Config struct {
	AccessKey  string
	SecretKey  string
	Region     string
	Bucket     string
	Endpoint   string
	DisableSSL bool
}

// Storage Configuration Variable
var storeS3Cfg storeS3Config

// Storage Connection Variable
var StoreS3 *session.Session

// Storage Connect Function
func storeS3Connect() *session.Session {
	var config *aws.Config

	// Create S3 Connection
	switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		config = &aws.Config{
			Credentials: credentials.NewStaticCredentials(storeS3Cfg.AccessKey, storeS3Cfg.SecretKey, ""),
			Region:      aws.String(storeS3Cfg.Region),
		}

	case "minio":
		// Set Endpoint URL Based On SSL Support
		var endpoint string
		if storeS3Cfg.DisableSSL {
			endpoint = "http://" + storeS3Cfg.Endpoint
		} else {
			endpoint = "https://" + storeS3Cfg.Endpoint
		}

		config = &aws.Config{
			Credentials:      credentials.NewStaticCredentials(storeS3Cfg.AccessKey, storeS3Cfg.SecretKey, ""),
			Endpoint:         aws.String(endpoint),
			Region:           aws.String(storeS3Cfg.Region),
			DisableSSL:       aws.Bool(storeS3Cfg.DisableSSL),
			S3ForcePathStyle: aws.Bool(true),
		}
	}

	// Return Session
	return session.New(config)
}

func StoreS3GetFileLink(fileName string) (string, error) {
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		// Return Composed URL Based on Storage Driver
		switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
		case "aws":
			return "https://s3." + storeS3Cfg.Region + ".amazonaws.com/" + storeS3Cfg.Bucket + "/" + fileName, nil
		case "minio":
			if storeS3Cfg.DisableSSL {
				return "http://" + storeS3Cfg.Endpoint + "/" + storeS3Cfg.Bucket + "/" + fileName, nil
			} else {
				return "https://" + storeS3Cfg.Endpoint + "/" + storeS3Cfg.Bucket + "/" + fileName, nil
			}
		}
	}

	return "", errors.New("No storage driver defined")
}

func StoreS3UploadFile(fileName string) error {
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		// Open File for Upload
		fileContent, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer fileContent.Close()

		// Create Uploader and Do Upload
		uploader := s3manager.NewUploader(StoreS3)
		result, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(storeS3Cfg.Bucket),
			Key:    aws.String(filepath.Base(fileName)),
			Body:   fileContent,
		})

		// Check for Upload Error
		if err != nil {
			return err
		}

		// Upload Success
		log.Println("File " + fileName + " successfully uploaded to " + result.Location)

		// Return No Error
		return nil
	}

	return errors.New("No storage driver defined")
}
