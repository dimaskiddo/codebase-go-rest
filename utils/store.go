package utils

import (
	"strings"
)

func initStore() {
	// Store Configuration Value
	switch strings.ToLower(Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		storeS3Cfg.AccessKey = Config.GetString("STORAGE_ACCESS_KEY")
		storeS3Cfg.SecretKey = Config.GetString("STORAGE_SECRET_KEY")
		storeS3Cfg.Region = Config.GetString("STORAGE_REGION")
		storeS3Cfg.Bucket = Config.GetString("STORAGE_BUCKET")

		StoreS3 = storeS3Connect()

	case "minio":
		storeS3Cfg.AccessKey = Config.GetString("STORAGE_ACCESS_KEY")
		storeS3Cfg.SecretKey = Config.GetString("STORAGE_SECRET_KEY")
		storeS3Cfg.Region = Config.GetString("STORAGE_REGION")
		storeS3Cfg.Bucket = Config.GetString("STORAGE_BUCKET")
		storeS3Cfg.Endpoint = Config.GetString("STORAGE_ENDPOINT")
		storeS3Cfg.DisableSSL = Config.GetBool("STORAGE_DISABLE_SSL")

		StoreS3 = storeS3Connect()
	}

}
