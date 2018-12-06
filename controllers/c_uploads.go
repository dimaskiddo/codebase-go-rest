package controllers

import (
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strings"

	"github.com/dimaskiddo/frame-go/utils"
)

// Function to Upload a File
func AddUpload(w http.ResponseWriter, r *http.Request) {
	// Limit Body Size with 1 MiB Margin
	r.Body = http.MaxBytesReader(w, r.Body, (utils.Config.GetInt64("SERVER_UPLOAD_LIMIT")+1)*int64(math.Pow(1024, 2)))

	// Get File Content from Multipart Data
	mpFile, mpHeader, err := r.FormFile("fileUpload")
	if err == nil {
		defer mpFile.Close()

		// Get File Metadata
		metaFileName := mpHeader.Filename
		metaFileSize := mpHeader.Size
		metaFileType := mpHeader.Header.Get("Content-Type")

		// Upload to Cloud Storage If Storage Driver Defined Else Save it to Local Storage
		switch strings.ToLower(utils.Config.GetString("STORAGE_DRIVER")) {
		case "aws", "minio":
			err := utils.StoreS3UploadFile(metaFileName, metaFileSize, metaFileType, mpFile)
			if err == nil {
				var response utils.Response

				// Set Response Data
				response.Status = true
				response.Code = http.StatusOK
				response.Message = "success"

				// Write Response Data to HTTP
				utils.ResponseWrite(w, response.Code, response)
			} else {
				utils.ResponseInternalError(w, err.Error())
				log.Println(err.Error())
			}
		default:
			// Default Save Uploaded File to Local Storage
			wrFile, err := os.OpenFile(utils.Config.GetString("SERVER_UPLOAD_PATH")+"/"+metaFileName, os.O_WRONLY|os.O_CREATE, 0666)
			if err == nil {
				defer wrFile.Close()
				var response utils.Response

				// Copy Uploaded File Data from Multipart Data
				io.Copy(wrFile, mpFile)

				// Set Response Data
				response.Status = true
				response.Code = http.StatusOK
				response.Message = "success"

				// Write Response Data to HTTP
				utils.ResponseWrite(w, response.Code, response)
			} else {
				utils.ResponseInternalError(w, err.Error())
				log.Println(err.Error())
			}
		}
	} else {
		utils.ResponseInternalError(w, err.Error())
		log.Println(err.Error())
	}
}
