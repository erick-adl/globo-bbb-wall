package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	// "github.com/aws/aws-lambda-go/lambda"
	. "github.com/erick-adl/globo-bbb-wall/lambda_total_hours/dao"
)

type TotalVotes struct {
	TotalVotes int `json:"total_votes"`
}

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Print("Getting all...\n")
	list, err := GetAllByHours()
	if err != nil {
		fmt.Print("Error GetAll...")
		fmt.Println(err)
		os.Exit(3)
	}

	fmt.Print("json.Marshal...\n")
	var x, _ = json.Marshal(list)
	fmt.Print(string(x))

	fmt.Print("Creating APIGatewayProxyResponse...\n")
	resp := events.APIGatewayProxyResponse{Headers: make(map[string]string)}
	resp.Body = string(x)
	resp.StatusCode = 200
	resp.Headers["Access-Control-Allow-Origin"] = "*"
	resp.Headers["Access-Control-Allow-Credentials"] = "true"
	fmt.Print("Finish...\n")
	fmt.Print("now, return\n")
	return resp, nil
}

func main() {

	lambda.Start(handleRequest)
}
