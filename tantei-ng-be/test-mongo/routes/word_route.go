package routes

import (
	"fmt"
	"test-mongo/controllers"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func WordRoutes(route string, data string, dbClient *mongo.Client) {
	var collection *mongo.Collection = dbClient.Database("tantei").Collection("sets")

	switch route {
	case "get_words":
		controllers.GetWords(collection)

	case "get_word_id":
		var object_id, err = bson.ObjectIDFromHex(data)

		if err != nil {
			fmt.Printf("Cannot convert data into bson.ObjectID!")
			return
		}

		controllers.GetWordById(collection, object_id)

	default:
		fmt.Printf("Route not found!")
	}
}
