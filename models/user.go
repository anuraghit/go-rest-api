package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id                    bson.ObjectId      `json:"id" bson:"_id"`
	Name                  string             `json:"name" bson:"name"`
	UserName              string             `json:"username" bson:"username"`
	UserDob               string             `json:"dob" bson:"dob"`
	UserAddress           Address            `json:"address" bson:"address"`
	UserDescription       string             `json:"description" bson:"description"`
	CreatedAt             string             `json:"createdAt" bson:"createdAt"`
}
