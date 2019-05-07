package dao

import (
	. "github.com/erick-adl/globo-bbb-wall/lambda_persistence/database"
	. "github.com/erick-adl/globo-bbb-wall/lambda_persistence/models"
	"gopkg.in/mgo.v2/bson"
)

var C_WALL = "wall"

func GetAll() ([]ListOfParticipant, error) {
	var listOfParticipant []ListOfParticipant
	err := DefaultDB().C(C_WALL).Find(bson.M{}).All(&listOfParticipant)
	return listOfParticipant, err
}

func GetByID(id string) (ListOfParticipant, error) {
	var listOfParticipant ListOfParticipant
	err := DefaultDB().C(C_WALL).FindId(bson.ObjectIdHex(id)).One(&listOfParticipant)
	return listOfParticipant, err
}

func GetByTime(time int) (ListOfParticipant, error) {
	var listOfParticipant ListOfParticipant
	err := DefaultDB().C(C_WALL).Find(bson.M{"time": time}).One(&listOfParticipant)
	return listOfParticipant, err
}

func Create(listOfParticipant ListOfParticipant) error {
	err := DefaultDB().C(C_WALL).Insert(&listOfParticipant)
	return err
}

func Update(id string, listOfParticipant ListOfParticipant) error {
	err := DefaultDB().C(C_WALL).UpdateId(bson.ObjectIdHex(id), &listOfParticipant)
	return err
}
