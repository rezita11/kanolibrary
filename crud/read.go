package crud

import (
	"context"
	"encoding/json"
	"kanolibrary/mongo"
	"kanolibrary/util"
	"kanolibrary/models"
	"go.mongodb.org/mongo-driver/bson"
)

type books struct {
	Name, Author, Publisher, Synopsis, Page string
}

func FindOne(collection string, query map[string]interface{}) ([]byte, error) {

	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	csr, err := db.Collection(collection).Find(ctx, query)
	util.ErrorChecker(err)
	defer csr.Close(ctx)

	result := make([]models.Books, 0)
	for csr.Next(ctx) {
		var row models.Books
		err := csr.Decode(&row)
		util.ErrorChecker(err)

		result = append(result, row)

	}
	data, err := json.Marshal(result)
	util.ErrorChecker(err)
	return data, err
}

func FindAll(collection string)([]byte, error) {
	ctx := context.Background()
	db, err := mongo.Connect()
	util.ErrorChecker(err)

	csr, err := db.Collection("books").Find(ctx, bson.D{{}})
	util.ErrorChecker(err)
	defer csr.Close(ctx)

	result := make([]books, 0)
	for csr.Next(ctx) {
		var row books
		err := csr.Decode(&row)
		util.ErrorChecker(err)

		result = append(result, row)
	}
	data, err := json.Marshal(result)
	util.ErrorChecker(err)
	return data, err
}
