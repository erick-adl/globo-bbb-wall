package models

import "gopkg.in/mgo.v2/bson"

type Participant struct {
	Name  string `bson:"name" json:"name"`
	Votes int    `bson:"votes" json:"votes"`
}

type ListOfParticipant struct {
	ID           bson.ObjectId `bson:"_id" json:"id"`
	Votes        int           `bson:"votes" json:"votes"`
	Time         int           `bson:"time" json:"time"`
	Participants []Participant `bson:"participants" json:"participants"`
}
