package main

import (
	"fmt"

	"test-mongo/controllers"
	"test-mongo/db"
	"test-mongo/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func main() {
	var err error
	var dbClient *mongo.Client = db.DbConnect()

	var wordCollection = models.WordCollection(dbClient)

	var test_id_str string = "6821d7285336e6824e98869b"
	test_id, err := bson.ObjectIDFromHex(test_id_str)
	// controllers.GetWordById(wordCollection, test_id)

	controllers.GetWordByIdP(wordCollection, test_id, "sets")

	if err != nil {
		fmt.Printf("Cannot create ObjectID from provided hex!\n")
	}
}
