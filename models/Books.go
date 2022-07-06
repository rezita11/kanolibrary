package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Books struct {
	Id			primitive.ObjectID `bson:"_id"`
	Name      	string `bson:"name"`
	Author     	string `bson: "author"`
	Publisher 	string `bson: "publisher"`
	Synopsis  	string `bson : "synopsis"`
	Page      	string `bson : "page"`
}
