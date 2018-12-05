package controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/dimaskiddo/frame-go/utils"
)

// Function to Upload a File
func AddUpload(w http.ResponseWriter, r *http.Request) {
	// Get File Content from Multipart Data
	fileUploadContent, fileUploadHeader, err := r.FormFile("fileUpload")
	if err == nil {
		defer fileUploadContent.Close()
		fileUploadName := fileUploadHeader.Filename

		// Create Uploaded File
		fileUploadWrite, err := os.OpenFile(os.Getenv("CONFIG_PATH")+"/uploads/"+fileUploadName, os.O_WRONLY|os.O_CREATE, 0666)
		if err == nil {
			var response utils.Response
			defer fileUploadWrite.Close()

			// Copy Uploaded File Data from Multipart Data
			io.Copy(fileUploadWrite, fileUploadContent)

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
	} else {
		utils.ResponseInternalError(w, err.Error())
		log.Println(err.Error())
	}
}
