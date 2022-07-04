package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	Id     primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title  string             `json:"title,omitempty" bson:"title"`
	Author string             `json:"author,omitempty" bson:"author"`
	Year   string             `json:"year,omitempty" bson:"year"`
}
