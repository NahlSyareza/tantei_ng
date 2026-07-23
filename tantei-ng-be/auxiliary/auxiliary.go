package auxiliary

import (
	"context"
	"fmt"
	"tantei-ng/db"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func Populate(test *bson.M, field_name string, ref_collection_name string, opts bson.D) {
	var dbClient = db.DbConnect()
	var ref_collection = dbClient.Database("tantei").Collection(ref_collection_name)
	var err error
	var end_result []bson.M

	// Still hardcoded:
	var selectedField bson.A = (*test)[field_name].(bson.A)

	for i := range selectedField {
		var result bson.M
		fmt.Printf("%s\n", selectedField[i])
		// Still hardcoded
		err = ref_collection.FindOne(context.TODO(), bson.D{{"_id", selectedField[i]}}, options.FindOne().SetProjection(opts)).Decode(&result)

		if err == mongo.ErrNoDocuments {
			fmt.Printf("Document not found!\n")
			continue
		}

		fmt.Printf("Document found!\n")

		end_result = append(end_result, result)
	}

	(*test)["sets"] = end_result
}
