package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var (
	// Set session and other variables globally
	// so they can be re-used
	AWSSession *session.Session
	AWSBucket  = os.Getenv("BUCKET_NAME")
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Verify request method
	if r.Method != "POST" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		resp := map[string]interface{}{
			"message": "method not allowed!",
		}
		json.NewEncoder(w).Encode(resp)
	}

	// Parse file from request
	err := r.ParseMultipartForm(0)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		resp := map[string]interface{}{
			"message": "failed to parse form request",
		}
		json.NewEncoder(w).Encode(resp)
	}

	// Get file data from parsed request
	fhs := r.MultipartForm.File["file"]
	for _, fh := range fhs {
		file, _, _ := r.FormFile("file")
		// Get the uploaded file name
		//fileName := fh.Filename

		// Setup uploader to AWS S3 bucket
		// with the session created at the beginning
		uploader := s3manager.NewUploader(AWSSession)

		// Prepare the upload payload
		payload := &s3manager.UploadInput{
			// Set which bucket to upload
			Bucket: &AWSBucket,
			// Set file name
			Key: &fh.Filename,
			// File to be uploaded
			Body: file,
		}

		result, err := uploader.Upload(payload)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			resp := map[string]interface{}{
				"message": "failed to upload file to S3 bucket",
			}
			json.NewEncoder(w).Encode(resp)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		resp := map[string]interface{}{
			"message":       "file uploaded",
			"file_location": result.Location,
		}
		json.NewEncoder(w).Encode(resp)
	}

}

func main() {
	log.Println("server starts!")
	fmt.Println("ini go-aws")

	// Setup AWS session
	// Automatically get needed configuration
	// from environment variables
	session, err := session.NewSession()
	if err != nil {
		log.Fatal(err)
	}

	AWSSession = session

	router := http.NewServeMux()
	router.HandleFunc("/upload", UploadFile)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()

}
