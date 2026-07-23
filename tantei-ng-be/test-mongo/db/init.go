package db

import (
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DbConnect() *mongo.Client {
	var uri string = "mongodb://admin:FullAutoWHEELDRIVERS%24%24%24%25%25%25@100.117.82.121:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.5.1"
	var docs string = "www.mongodb.com/docs/drivers/go/current/"
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		// Perhaps something better than panic is better
		panic(err)
		// return nil
	}

	return client
}
