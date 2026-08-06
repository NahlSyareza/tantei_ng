package controllers

import (
	"context"
	"net/http"
	"tantei-ng/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetNgSets(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	var collection *mongo.Collection = models.NgSetCollection()

	var opts = options.Find().SetProjection(bson.D{{"items", 0}})
	var cursor, err = collection.Find(context.TODO(), bson.D{}, opts)

	var results []models.NgSetSchema

	err = cursor.All(context.TODO(), &results)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, results)
}

func GetNgSet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var collection *mongo.Collection = models.NgSetCollection()
	var oid = c.Param("oid")

	var _id, err = bson.ObjectIDFromHex(oid)

	if err != nil {
		panic(err)
	}

	var filter bson.D = bson.D{{"_id", _id}}

	var result models.NgSetSchema
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, result)
}

func CreateNgSet(c *gin.Context) {
	var collection *mongo.Collection = models.NgSetCollection()

	var doc models.NgSetSchema

	var err = c.BindJSON(&doc)

	// doc.Id = bson.NewObjectID()
	// doc.CreatedAt = bson.NewObjectID().Timestamp()
	// doc.UpdatedAt = bson.NewObjectID().Timestamp()

	if err != nil {
		panic(err)
	}

	collection.InsertOne(context.TODO(), doc)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Created new set!",
	})
}

func AddNgSetItem(c *gin.Context) {
	var collection *mongo.Collection = models.NgSetCollection()
	var oid = c.Param("oid")

	var _id, err = bson.ObjectIDFromHex(oid)

	var result models.NgSetSchema
	var doc models.NgWordSchema

	err = c.BindJSON(&doc)

	if err != nil {
		panic(err)
	}

	var opts = options.FindOneAndUpdate().SetReturnDocument(options.After)
	var filter bson.D = bson.D{{"_id", _id}}
	var update bson.D = bson.D{{"$push", bson.D{{"items", doc}}}}

	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, result)
}

func RemoveNgSetItem(c *gin.Context) {
	var collection *mongo.Collection = models.NgSetCollection()
	var oid = c.Param("oid")

	var _id, err = bson.ObjectIDFromHex(oid)
	var doc models.NgWordSchema

	err = c.BindJSON(&doc)

	if err != nil {
		panic(err)
	}

	var opts = options.FindOneAndUpdate().SetReturnDocument(options.After)
	var filter = bson.D{{"_id", _id}}
	var update = bson.D{{"$pull", bson.D{{"items", doc}}}}

	var result models.NgSetSchema
	err = collection.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&result)

	c.IndentedJSON(http.StatusOK, result)
}

func AddManyNgSetItem(c *gin.Context) {
	var collection mongo.Collection = *models.NgSetCollection()
	var oid = c.Param("oid")

	var docs []models.NgWordSchema

	var err = c.BindJSON(&docs)

	if err != nil {
		panic(err)
	}

	_id, err := bson.ObjectIDFromHex(oid)

	if err != nil {
		panic(err)
	}

	var result models.NgSetSchema

	var options = options.FindOneAndUpdate().SetReturnDocument(options.After)
	var filter bson.D = bson.D{{"_id", _id}}

	for _, e := range docs {
		var update bson.D = bson.D{{"$push", bson.D{{"items", e}}}}
		err = collection.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&result)

		if err != nil {
			panic(err)
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}

func RemoveManyNgSetItem(c *gin.Context) {
	var collection mongo.Collection = *models.NgSetCollection()
	var oid = c.Param("oid")

	var docs []models.NgWordSchema

	var err = c.BindJSON(&docs)

	if err != nil {
		panic(err)
	}

	_id, err := bson.ObjectIDFromHex(oid)

	if err != nil {
		panic(err)
	}

	var result models.NgSetSchema

	var options = options.FindOneAndUpdate().SetReturnDocument(options.After)
	var filter bson.D = bson.D{{"_id", _id}}

	for _, e := range docs {
		var update bson.D = bson.D{{"$pull", bson.D{{"items", e}}}}
		err = collection.FindOneAndUpdate(context.TODO(), filter, update, options).Decode(&result)

		if err != nil {
			panic(err)
		}
	}

	c.IndentedJSON(http.StatusOK, result)
}
