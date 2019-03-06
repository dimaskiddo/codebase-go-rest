package controller

import (
	"io"
	"math"
	"net/http"
	"os"
	"strings"

	svc "github.com/dimaskiddo/frame-go/service"
)

// UploadFile Function to Upload a File
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Limit Body Size with 1 MiB Margin
	err := r.ParseMultipartForm((svc.Config.GetInt64("SERVER_UPLOAD_LIMIT") + 1) * int64(math.Pow(1024, 2)))
	if err != nil {
		svc.ResponseInternalError(w, err.Error())
		return
	}

	// Get File Content from Multipart Data
	mpFile, mpHeader, err := r.FormFile("file")
	if err != nil {
		svc.ResponseBadRequest(w, err.Error())
		return
	}
	defer mpFile.Close()

	// Get File Metadata
	metaFileName := mpHeader.Filename
	metaFileSize := mpHeader.Size
	metaFileType := mpHeader.Header.Get("Content-Type")

	// Upload to Cloud Storage If Storage Driver Defined Else Save it to Local Storage
	switch strings.ToLower(svc.Config.GetString("STORAGE_DRIVER")) {
	case "aws", "minio":
		err := svc.StoreS3UploadFile(metaFileName, metaFileSize, metaFileType, mpFile)
		if err != nil {
			svc.ResponseInternalError(w, err.Error())
			return
		}

		svc.ResponseSuccess(w, "")
	default:
		// Default Save Uploaded File to Local Storage
		wrFile, err := os.OpenFile(svc.Config.GetString("SERVER_UPLOAD_PATH")+"/"+metaFileName, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			svc.ResponseInternalError(w, err.Error())
			return
		}
		defer wrFile.Close()

		// Copy Uploaded File Data from Multipart Data
		io.Copy(wrFile, mpFile)

		svc.ResponseSuccess(w, "")
	}
}
