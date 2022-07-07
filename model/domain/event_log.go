package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type EventLog struct {
	Id   primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Type string             `json:"tupe,omitempty" bson:"type"`
	Code string             `json:"code" bson:"code"`
}
