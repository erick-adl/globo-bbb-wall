package dao

import (
	. "github.com/erick-adl/globo-bbb-wall/lambda_update_ssm/database"
	. "github.com/erick-adl/globo-bbb-wall/lambda_update_ssm/models"
	"gopkg.in/mgo.v2/bson"
)

var C_WALL = "wall"

func GetAll() ([]Participant, error) {
	var listOfParticipant []ListOfParticipant
	var list []Participant
	err := DefaultDB().C(C_WALL).Find(bson.M{}).All(&listOfParticipant)
	if err != nil {
		return list, err
	}
	list = make([]Participant, len(listOfParticipant[0].Participants))
	for x := range listOfParticipant {
		for y := range listOfParticipant[x].Participants {
			list[y].Name = listOfParticipant[x].Participants[y].Name
			list[y].Votes += listOfParticipant[x].Participants[y].Votes
		}
	}
	return list, err
}
