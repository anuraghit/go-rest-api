package models

type Address struct {
	City     string  `json:"city" bson:"city"`
	Country  string  `json:"country" bson:"country"`
	StreatNo string  `json:"streatNo" bson:"streatNo"`
	Pincode  string  `json:"pincode" bson:"pincode"`
	Lat      float32 `json:"lat" bson:"lat"`
	Lon      float32 `json:"lon" bson:"lon"`
}
