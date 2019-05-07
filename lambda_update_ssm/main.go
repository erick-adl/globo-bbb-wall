package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	. "github.com/erick-adl/globo-bbb-wall/lambda_update_ssm/dao"
)

var svc *ssm.SSM

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, e events.CloudWatchEvent) {

	fmt.Println("Configuring...")
	err := Configure()

	if err != nil {
		fmt.Print("Error configuring...")
		fmt.Println(err)
		os.Exit(2)
	}

	participants, err := GetAll()
	if err != nil {
		fmt.Print("Error GetAll...")
		fmt.Println(err)
		os.Exit(3)
	}
	if len(participants) == 0 {
		fmt.Print("Participants not found...")
		os.Exit(4)
	}
	var participants_name_parameter = "participants_name"
	var participants_value_parameter []string
	participants_value_parameter = nil
	for i := range participants {
		CreateParameter(participants[i].Name, strconv.Itoa(participants[i].Votes))
		participants_value_parameter = append(participants_value_parameter, participants[i].Name)
	}
	CreateParameterMultString(participants_name_parameter, participants_value_parameter)

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

	fmt.Print(*resp.Parameters[0].Value)
	return *resp.Parameters[0].Value, err
}

func CreateParameter(name string, value string) (bool, error) {
	putParameterInput := &ssm.PutParameterInput{
		Name:      aws.String(name),
		Value:     aws.String(value),
		Type:      aws.String("String"),
		Overwrite: aws.Bool(true),
	}

	// This API call returns an empty struct
	resp, err := svc.PutParameter(putParameterInput)
	if resp != nil {
		return true, err
	}
	return false, err
}

func CreateParameterMultString(name string, value []string) (bool, error) {
	putParameterInput := &ssm.PutParameterInput{
		Name:      aws.String(name),
		Value:     aws.String(strings.Join(value, ",")),
		Type:      aws.String("StringList"),
		Overwrite: aws.Bool(true),
	}

	// This API call returns an empty struct
	resp, err := svc.PutParameter(putParameterInput)
	if resp != nil {
		return true, err
	}
	return false, err
}
