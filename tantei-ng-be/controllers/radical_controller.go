package controllers

import (
	"context"
	"net/http"
	"tantei-ng/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func GetRadicalLists(c *gin.Context) {
	collection := models.RadicalListCollection()

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		panic(err)
	}

	var results []models.RadicalListSchema

	err = cursor.All(context.TODO(), &results)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, results)
}

func CreateRadicalList(c *gin.Context) {
	collection := models.RadicalListCollection()

	var doc models.RadicalListSchema

	err := c.BindJSON(&doc)

	if err != nil {
		panic(err)
	}

	collection.InsertOne(context.TODO(), doc)

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Inserted new Radical List!",
	})
}

func AddItemsRadicalList(c *gin.Context) {
	collection := models.RadicalListCollection()

	radicalList := c.Param("radical_list")
	radicalListObjectID, err := bson.ObjectIDFromHex(radicalList)

	if err != nil {
		panic(err)
	}

	var doc []string

	err = c.BindJSON(&doc)

	if err != nil {
		panic(err)
	}

	filter := bson.D{{"_id", radicalListObjectID}}
	update := bson.D{{"$push", bson.D{{"radical_variants", bson.D{{"$each", doc}}}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "Successfully added new items"})
}
