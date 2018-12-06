package utils

import (
	"errors"
	"log"
	"path/filepath"
	"strings"

	minio "github.com/minio/minio-go"
)

// Storage Configuration Struct
type storeS3Config struct {
	UseSSL    bool
	Endpoint  string
	AccessKey string
	SecretKey string
	Region    string
	Bucket    string
}

// Storage Configuration Variable
var storeS3Cfg storeS3Config

// Storage Connection Variable
var StoreS3 *minio.Client

// Storage Connect Function
func storeS3Connect() *minio.Client {
	// Get Storage Connection
	switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		client, err := minio.New("s3.amazonaws.com", storeS3Cfg.AccessKey, storeS3Cfg.SecretKey, storeS3Cfg.UseSSL)
		if err != nil {
			log.Fatalln(err)
		}
		return client
	case "minio":
		client, err := minio.New(storeS3Cfg.Endpoint, storeS3Cfg.AccessKey, storeS3Cfg.SecretKey, storeS3Cfg.UseSSL)
		if err != nil {
			log.Fatalln(err)
		}
		return client
	}

	return nil
}

func StoreS3UploadFile(fileName string, contentType string) error {
	// Check If Storage Driver Declared
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		// Check If Bucket Exists
		bucketExists, err := StoreS3.BucketExists(storeS3Cfg.Bucket)
		if err != nil {
			return err
		} else {
			if !bucketExists {
				// If Bucket Not Exists Then Create Bucket
				err := StoreS3.MakeBucket(storeS3Cfg.Bucket, storeS3Cfg.Region)
				if err != nil {
					return err
				}
			} else {
				// If Bucket Exists Then Try to Upload File
				_, err := StoreS3.FPutObject(storeS3Cfg.Bucket, filepath.Base(fileName), fileName, minio.PutObjectOptions{ContentType: contentType})
				if err != nil {
					return err
				}

				log.Println("Successfully uploaded " + filepath.Base(fileName))
				return nil
			}
		}
	}

	// Default Return
	return errors.New("No storage driver defined")
}

func StoreS3GetFileLink(fileName string) (string, error) {
	// Check If Storage Driver Declared
	if len(strings.ToLower(Config.GetString("STORAGE_DRIVER"))) != 0 {
		// Return Composed URL Based on Storage Driver
		switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
		case "aws":
			return "https://s3.amazonaws.com/" + storeS3Cfg.Bucket + "/" + fileName, nil
		case "minio":
			if storeS3Cfg.UseSSL {
				return "https://" + storeS3Cfg.Endpoint + "/" + storeS3Cfg.Bucket + "/" + fileName, nil
			} else {
				return "http://" + storeS3Cfg.Endpoint + "/" + storeS3Cfg.Bucket + "/" + fileName, nil
			}
		}
	}

	// Default Return
	return "", errors.New("No storage driver defined")
}
