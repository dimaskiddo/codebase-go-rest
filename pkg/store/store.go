package store

import (
	"strings"

	"github.com/dimaskiddo/codebase-go-rest/pkg/server"
)

// Initialize Function in Store
func init() {
	// Store Configuration Value
	switch strings.ToLower(server.Config.GetString("STORAGE_DRIVER")) {
	case "aws":
		s3Cfg.AccessKey = server.Config.GetString("STORAGE_ACCESS_KEY")
		s3Cfg.SecretKey = server.Config.GetString("STORAGE_SECRET_KEY")
		s3Cfg.Region = server.Config.GetString("STORAGE_REGION")
		s3Cfg.Bucket = server.Config.GetString("STORAGE_BUCKET")

		s3 = s3Connect()
	case "minio":
		s3Cfg.UseSSL = server.Config.GetBool("STORAGE_USE_SSL")
		s3Cfg.Endpoint = server.Config.GetString("STORAGE_ENDPOINT")
		s3Cfg.AccessKey = server.Config.GetString("STORAGE_ACCESS_KEY")
		s3Cfg.SecretKey = server.Config.GetString("STORAGE_SECRET_KEY")
		s3Cfg.Region = server.Config.GetString("STORAGE_REGION")
		s3Cfg.Bucket = server.Config.GetString("STORAGE_BUCKET")

		s3 = s3Connect()
	}
}
