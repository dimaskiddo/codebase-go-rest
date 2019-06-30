package store

import (
	"strings"

	"github.com/dimaskiddo/codebase-go-rest/hlp"
)

// Initialize Function in Store
func init() {
	// Store Configuration Value
	switch strings.ToLower(hlp.Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		s3Cfg.AccessKey = hlp.Config.GetString("STORAGE_ACCESS_KEY")
		s3Cfg.SecretKey = hlp.Config.GetString("STORAGE_SECRET_KEY")
		s3Cfg.Region = hlp.Config.GetString("STORAGE_REGION")
		s3Cfg.Bucket = hlp.Config.GetString("STORAGE_BUCKET")

		s3 = s3Connect()
	case "minio":
		s3Cfg.UseSSL = hlp.Config.GetBool("STORAGE_USE_SSL")
		s3Cfg.Endpoint = hlp.Config.GetString("STORAGE_ENDPOINT")
		s3Cfg.AccessKey = hlp.Config.GetString("STORAGE_ACCESS_KEY")
		s3Cfg.SecretKey = hlp.Config.GetString("STORAGE_SECRET_KEY")
		s3Cfg.Region = hlp.Config.GetString("STORAGE_REGION")
		s3Cfg.Bucket = hlp.Config.GetString("STORAGE_BUCKET")

		s3 = s3Connect()
	}
}
