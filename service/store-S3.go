package service

import (
	"errors"
	"log"
	"mime/multipart"
	"strings"

	minio "github.com/minio/minio-go"
)

// StoreS3 Configuration Struct
type storeS3Config struct {
	UseSSL    bool
	Endpoint  string
	AccessKey string
	SecretKey string
	Region    string
	Bucket    string
}

// StoreS3 Configuration Variable
var storeS3Cfg storeS3Config

// StoreS3 Variable
var StoreS3 *minio.Client

// StoreS3 Connect Function
func storeS3Connect() *minio.Client {
	switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		conn, err := minio.New("s3.amazonaws.com", storeS3Cfg.AccessKey, storeS3Cfg.SecretKey, true)
		if err != nil {
			log.Fatalln(err)
		}
		return conn
	case "minio":
		conn, err := minio.New(storeS3Cfg.Endpoint, storeS3Cfg.AccessKey, storeS3Cfg.SecretKey, storeS3Cfg.UseSSL)
		if err != nil {
			log.Fatalln(err)
		}
		return conn
	default:
		return nil
	}
}

// StoreS3UploadFile Function to Upload File to S3 Storage
func StoreS3UploadFile(fileName string, fileSize int64, fileType string, fileStream multipart.File) error {
	switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
	case "aws", "minio":
		// Check If Bucket Exists
		bucketExists, err := StoreS3.BucketExists(storeS3Cfg.Bucket)
		if err != nil {
			return err
		}

		// If Bucket Not Exists Then Create Bucket
		if !bucketExists {
			err := StoreS3.MakeBucket(storeS3Cfg.Bucket, storeS3Cfg.Region)
			if err != nil {
				return err
			}
		}

		// Try to Upload File into Bucket
		n, err := StoreS3.PutObject(storeS3Cfg.Bucket, fileName, fileStream, fileSize, minio.PutObjectOptions{ContentType: fileType})
		if err != nil {
			return err
		}

		log.Printf("Successfully uploaded '%s', with size %d\n", fileName, n)
		return nil
	default:
		return errors.New("No storage driver defined")
	}
}

// StoreS3GetFileLink Function to Get Link for Uploaded File in S3 Storage
func StoreS3GetFileLink(fileName string) (string, error) {
	switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		return "https://s3-" + storeS3Cfg.Region + ".amazonaws.com/" + storeS3Cfg.Bucket + "/" + strings.Replace(fileName, " ", "+", -1), nil
	case "minio":
		if !storeS3Cfg.UseSSL {
			return "http://" + storeS3Cfg.Endpoint + "/" + storeS3Cfg.Bucket + "/" + fileName, nil
		}
		return "https://" + storeS3Cfg.Endpoint + "/" + storeS3Cfg.Bucket + "/" + fileName, nil
	default:
		return "", errors.New("No storage driver defined")
	}
}
