package models

import "gopkg.in/mgo.v2/bson"

type Property struct {
	PropID       bson.ObjectId `json:"id" bson:"_id"`
	PropertyName string        `json:"propertyname" bson:"propertyname"`
	Address      string        `json:"address" bson:"address"`
	City         string        `json"city" bson:"city"`
	Bedrooms     int32         `json"bedrooms" bson:"bedrooms"`
}
