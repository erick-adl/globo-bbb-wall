package dao

import (
	. "github.com/erick-adl/globo-bbb-wall/lambda_total_hours/database"
	. "github.com/erick-adl/globo-bbb-wall/lambda_total_hours/models"
	"gopkg.in/mgo.v2/bson"
)

var C_WALL = "wall"

func GetAllByHours() ([]HourByHour, error) {

	var hourByHour []HourByHour
	var listOfParticipant []ListOfParticipant
	err := DefaultDB().C(C_WALL).Find(bson.M{}).All(&listOfParticipant)
	hourByHour = make([]HourByHour, len(listOfParticipant))
	for i := range listOfParticipant {
		hourByHour[i].Hour = listOfParticipant[i].Time
		hourByHour[i].Participants = listOfParticipant[i].Participants
	}
	return hourByHour, err

}
