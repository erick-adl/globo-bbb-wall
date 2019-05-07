package main

import (
	"context"
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2/bson"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	. "github.com/erick-adl/globo-bbb-wall/lambda_persistence/dao"
	. "github.com/erick-adl/globo-bbb-wall/lambda_persistence/models"
)

func init() {
	fmt.Println("Getting message from sqs...")
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)

		if len(message.Body) == 0 {
			fmt.Println("Sqs empty message...")
		} else {
			bytes := []byte(message.Body)
			var sqsMsgDoc = ListOfParticipant{}
			err := json.Unmarshal(bytes, &sqsMsgDoc)
			if err != nil {
				panic(err)
			} else {
				fmt.Println("Searching...")
				listFound, err := GetAll()
				if err != nil {
					fmt.Print(err)
					panic(err)
				}
				if len(listFound) == 0 {
					fmt.Print("ok")
					sqsMsgDoc.ID = bson.NewObjectId()
					err = Create(sqsMsgDoc)
					if err != nil {
						fmt.Print(err)
						panic(err)
					}
					fmt.Print("ok")
				} else {

					docFound, err := GetByTime(sqsMsgDoc.Time)
					if err != nil {
						fmt.Print(err)
					}
					if docFound.ID != "" {
						for i := range docFound.Participants {
							docFound.Participants[i].Votes += sqsMsgDoc.Participants[i].Votes
							docFound.Votes += sqsMsgDoc.Participants[i].Votes
						}

						err = Update(docFound.ID.Hex(), docFound)
						if err != nil {
							fmt.Print(err)
							panic(err)
						}
					} else {
						fmt.Print("ok")
						sqsMsgDoc.ID = bson.NewObjectId()
						err = Create(sqsMsgDoc)
						if err != nil {
							fmt.Print(err)
							panic(err)
						}
						fmt.Print("ok")
					}
				}
			}
		}
	}
	fmt.Println("finish...")

	return nil
}

func main() {

	lambda.Start(handler)
}
