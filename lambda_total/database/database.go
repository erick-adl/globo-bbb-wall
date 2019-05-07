package database

import (
	"crypto/tls"
	"log"
	"net"
	"os"

	mgo "gopkg.in/mgo.v2"
)

var (
	dbPrefix    = "bbb"
	Environment = os.Getenv("MONGODB_ENV")
)

var mongodbSession *mgo.Session

func DBSession() *mgo.Session {
	if mongodbSession != nil {
		return mongodbSession
	}

	tlsConfig := &tls.Config{}

	dialInfo := &mgo.DialInfo{
		Addrs:    []string{""},
		Database: "",
		Username: "",
		Password: "",
	}
	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	mongodbSession, err := mgo.DialWithInfo(dialInfo)
	if mongodbSession == nil || err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
	}

	mongodbSession.SetSafe(&mgo.Safe{})
	return mongodbSession
}

func DB(name string) *mgo.Database {
	return DBSession().DB(name)
}
func DefaultDB() *mgo.Database {
	switch Environment {
	case "test":
		{
			return DB(dbPrefix + "-test")
		}
	case "production":
		{
			return DB(dbPrefix + "-production")
		}
	}

	return DB(dbPrefix + "-development")
}
