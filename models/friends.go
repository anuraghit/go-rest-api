package models

import "gopkg.in/mgo.v2/bson"

type Friend struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	Following bool          `json:"following" bson:"following"`
	Follower  bool          `json:"follower" bson:"follower"`
}

type FollowingRequest struct {
	UserId bson.ObjectId `json:"userid" bson:"userid"`
}

type FollowRequest struct {
	UserId bson.ObjectId `json:"userid" bson:"userid"`
}
