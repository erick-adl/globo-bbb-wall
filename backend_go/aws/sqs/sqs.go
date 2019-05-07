package sqs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

var svc *sqs.SQS
var queue_url string

func getQueueSqs() string {
	url := os.Getenv("SQS_QUEUE")

	if url == "" {
		url = "teste"
	}

	return url
}

func getAwsRegion() string {
	region := os.Getenv("AWS_REGION")

	if region == "" {
		region = "us-east-1"
	}

	return region
}

func Configure() (int, error) {

	region := getAwsRegion()
	queue_sqs := getQueueSqs()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return -1, err
	}

	svc = sqs.New(sess)
	result, err := svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: aws.String(queue_sqs),
	})

	if err != nil {
		return -1, err
	}

	queue_url = *result.QueueUrl
	return 0, err
}

func getClinet() *sqs.SQS {
	return svc
}

func SendMessage(input string) (*sqs.SendMessageOutput, error) {

	result, err := getClinet().SendMessage(&sqs.SendMessageInput{
		DelaySeconds: aws.Int64(10),
		MessageBody:  aws.String(input),
		QueueUrl:     &queue_url,
	})

	if err != nil {
		fmt.Println("Error", err)
		return result, err
	}
	// fmt.Println("Success", *result.MessageId)
	return result, err

}
