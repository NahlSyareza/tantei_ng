package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func DbConnect() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found!")
	}

	var uri string = os.Getenv("MONGODB_URI")
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
	}

	return client
}
