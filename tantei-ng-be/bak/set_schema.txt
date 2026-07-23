package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SetSchema struct {
	Id        bson.ObjectID `bson:"_id"`
	Name      string        `bson:"name"`
	Owner     bson.ObjectID `bson:"owner"`
	T         string        `bson:"__t"`
	Items     []any         `bson:"items"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

func SetCollection(dbClient *mongo.Client) *mongo.Collection {
	var collection = dbClient.Database("tantei").Collection("sets")

	return collection
}
