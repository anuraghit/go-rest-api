package database

import "gopkg.in/mgo.v2"

var MongoDB *mgo.Session

func SetupMongo() {
	s, err := mgo.Dial("mongodb://127.0.0.1:27017")
	if err != nil {
		panic(err.Error())
	}
	MongoDB = s
}
