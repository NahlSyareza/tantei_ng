package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"test-mongo/auxiliary"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func GetWords(collection *mongo.Collection) {
	var cursor, err = collection.Find(context.TODO(), bson.D{})

	var results []bson.M

	err = cursor.All(context.TODO(), &results)

	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(results)

	fmt.Printf("%s\n", jsonData)
}

func GetWordById(collection *mongo.Collection, id bson.ObjectID) {
	var result bson.M
	var err error = collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&result)

	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(result, "", "	")

	fmt.Printf("%s\n", jsonData)
}

func GetWordByIdP(collection *mongo.Collection, id bson.ObjectID, ref_collection_name string) {
	var result bson.M
	var err error = collection.FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(&result)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	auxiliary.Populate(&result, "sets", "sets", bson.D{{"name", 1}, {"owner", 1}})

	// fmt.Printf("%s\n", *test_pop)

	jsonData, err := json.MarshalIndent(result, "", "	")

	fmt.Printf("%s\n", jsonData)
}
