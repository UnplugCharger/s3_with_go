package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	fmt.Println("Downloads and Uploads")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"),
	})
	if err != nil {
		log.Fatal("Could not get session")
	}
	uploadItem(sess)
}


func uploadItem(sess *session.Session) {
	f, err := os.Open("z.txt")
	if err != nil {
		log.Fatal("Could not open file")
	}
	defer f.Close()

	uploader := s3manager.NewUploader(sess)
	result, err := uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("private"),
		Bucket: aws.String("bidmania-upload-trial"),
		Key:    aws.String("z.txt"),
		Body:   f,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Upload Result: %+v\n", result)
}