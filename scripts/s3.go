package main

import (
	"encoding/json"
    "fmt"
    "io/ioutil"
    "log"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type S3Bucket struct {
    Bucket string `json:"bucket"`
    Key    string `json:"key"`
}

var pageNum int = 0
var s3Buckets []S3Bucket
var sess *session.Session

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

}