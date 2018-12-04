package utils

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
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

// Storage Path Variable
var StorageS3Path string

// Storage Connection Variable
var StorageS3 *s3.S3

// Storage Connect Function
func initStorageS3() {
	region := aws.String(storageS3Cfg.Region)

	switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		config := &aws.Config{
			Credentials: credentials.NewStaticCredentials(storageS3Cfg.AccessKey, storageS3Cfg.SecretKey, ""),
			Region:      region,
		}
	case "minio":
		config := &aws.Config{
			Credentials:      credentials.NewStaticCredentials(storageS3Cfg.AccessKey, storageS3Cfg.SecretKey, ""),
			Endpoint:         aws.String(storageS3Cfg.Endpoint),
			Region:           region,
			DisableSSL:       aws.Bool(storageS3Cfg.DisableSSL),
			S3ForcePathStyle: aws.Bool(true),
		}
	}

	StorageS3 = s3.New(session.New(config))
	StorageS3Path = storageS3Cfg.Endpoint + "/" + storageS3Cfg.Bucket
}

func StorageS3Upload() {

}
