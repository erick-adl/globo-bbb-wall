package dao

import (
	. "github.com/erick-adl/globo-bbb-wall/lambda_update_ssm/database"
	. "github.com/erick-adl/globo-bbb-wall/lambda_update_ssm/models"
	"gopkg.in/mgo.v2/bson"
)

var C_WALL = "wall"

func GetTotal() (int, error) {
	var listOfParticipant []ListOfParticipant
	var total int
	err := DefaultDB().C(C_WALL).Find(bson.M{}).All(&listOfParticipant)
	if err != nil {
		return total, err
	}
	for x := range listOfParticipant {
		total += listOfParticipant[x].Votes
	}
	return total, err
}
