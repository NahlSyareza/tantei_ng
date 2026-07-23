package models

import (
	"tantei-ng/db"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type NgWordSchema struct {
	Kanji     string `bson:"kanji" json:"kanji"`
	Furigana  string `bson:"furigana" json:"furigana"`
	Latin     string `bson:"latin" json:"latin"`
	English   string `bson:"english" json:"english"`
	Indonesian string `bson:"indonesian" json:"indonesian"`
}

type NgSetSchema struct {
	Id        bson.ObjectID  `bson:"_id" json:"_id"`
	Name      string         `bson:"name" json:"name"`
	Items     []NgWordSchema `bson:"items,omitempty" json:"items,omitempty"`
	CreatedAt time.Time      `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time      `bson:"updatedAt" json:"updatedAt"`
}

func NgSetCollection() *mongo.Collection {
	var dbClient *mongo.Client = db.DbConnect()
	var collection *mongo.Collection = dbClient.Database("tantei").Collection("ng_sets")

	return collection
}

