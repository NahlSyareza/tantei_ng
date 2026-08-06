package models

import (
	"tantei-ng/db"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type NgWordSchema struct {
	Kanji      string `bson:"kanji" json:"kanji"`
	Furigana   string `bson:"furigana" json:"furigana"`
	Latin      string `bson:"latin" json:"latin"`
	English    string `bson:"english" json:"english"`
	Indonesian string `bson:"indonesian" json:"indonesian"`
}

type NgSetSchema struct {
	Id    bson.ObjectID  `bson:"_id,omitempty" json:"_id,omitempty"`
	Owner bson.ObjectID  `bson:"owner,omitempty" json:"owner,omitempty"`
	Name  string         `bson:"name,omitempty" json:"name,omitempty"`
	Items []NgWordSchema `bson:"items,omitempty" json:"items,omitempty"`
	// CreatedAt time.Time      `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	// UpdatedAt time.Time      `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type AccountSchema struct {
	Id       bson.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Name     string        `bson:"name,omitempty" json:"name,omitempty"`
	Email    string        `bson:"email,omitempty" json:"email,omitempty"`
	Password string        `bson:"password,omitempty" json:"password,omitempty"`
}

type TrackerSchema struct {
	Id            bson.ObjectID   `bson:"_id,omitempty" json:"_id,omitempty"`
	Owner         bson.ObjectID   `bson:"owner,omitempty" json:"owner,omitempty"`
	RemainingSets []bson.ObjectID `bson:"remaining_sets,omitempty" json:"remaining_sets,omitempty"`
	LastTime      time.Time       `bson:"last_time,omitempty" json:"last_time,omitempty"`
}

func NgSetCollection() *mongo.Collection {
	var dbClient *mongo.Client = db.DbConnect()
	var collection *mongo.Collection = dbClient.Database("tantei").Collection("ng_sets")

	return collection
}

func AccountCollection() *mongo.Collection {
	var dbClient *mongo.Client = db.DbConnect()
	var collection *mongo.Collection = dbClient.Database("tantei").Collection("accounts")

	return collection
}

func TrackerCollection() *mongo.Collection {
	var dbClient *mongo.Client = db.DbConnect()
	var collection *mongo.Collection = dbClient.Database("tantei").Collection("trackers")

	return collection
}
