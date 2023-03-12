package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BlogPost struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
}
