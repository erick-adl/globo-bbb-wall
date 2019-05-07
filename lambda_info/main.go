package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

var (
	svc *ssm.SSM
)

type ParticipantAnswerList struct {
	Participants []ParticipantAnswer `json:"participants"`
}

type ParticipantAnswer struct {
	Name  string `json:"Name"`
	Votes string `json:"Votes"`
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	err := Configure()

	if err != nil {
		fmt.Print("Error configuring...")
		fmt.Println(err)

	}
	var participantsNameParameter = "participants_name"
	var participantAnswerList = ParticipantAnswerList{}
	participantsListString, err := GetParameter(participantsNameParameter)
	participantsList := strings.Split(participantsListString, ",")
	for i := range participantsList {
		var temp ParticipantAnswer
		temp.Name = participantsList[i]
		value, err := GetParameter(participantsList[i])
		if err != nil {
			fmt.Print("Error GetParameter...")
			fmt.Println(err)

		}
		temp.Votes = value
		participantAnswerList.Participants = append(participantAnswerList.Participants, temp)
	}
	fmt.Println(participantAnswerList.Participants)

	var x, _ = json.Marshal(participantAnswerList)
	fmt.Print(string(x))
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Body = string(x)
	resp.StatusCode = 200
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Access-Control-Allow-Credentials"] = "true"

	return resp, nil
}

func main() {
	lambda.Start(handleRequest)
}

func Configure() error {

	region := "us-east-1"

	fmt.Println("Getting session...")
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		fmt.Print("Error creating sess...")
		return err
	}

	fmt.Println("Getting svc...")
	svc = ssm.New(sess, &aws.Config{
		MaxRetries: aws.Int(10),
		Region:     aws.String(region),
	})
	if err != nil {
		fmt.Print("Error creating svc...")
		return err
	}
	_, err = GetParameter("teste")
	if err != nil {
		return err
	}
	return nil
}

func GetParameter(p string) (string, error) {
	getParametersInput := &ssm.GetParametersInput{
		Names: []*string{aws.String(p)},
	}
	resp, err := svc.GetParameters(getParametersInput)
	if err != nil {
		fmt.Print(err)
		return "", err
	}

	return *resp.Parameters[0].Value, err
}
