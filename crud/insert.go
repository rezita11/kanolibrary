package crud

import (
	"context"
	"kanolibrary/models"
	"kanolibrary/mongo"
	"kanolibrary/util"
)

func Insert(collection string, query models.Books)(string, error) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	_,err = db.Collection(collection).InsertOne(ctx, query)
	util.ErrorChecker(err)

	return "Insert success", err
}