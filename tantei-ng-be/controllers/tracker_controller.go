package controllers

import (
	"context"
	"net/http"
	"tantei-ng/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type trackerHelper struct {
	SetId   string `json:"set_id,omitempty"`
	OwnerId string `json:"owner_id,omitempty"`
}

func RemoveItemTracker(c *gin.Context) {
	var helperDoc trackerHelper

	c.BindJSON(&helperDoc)

	setId, err := bson.ObjectIDFromHex(helperDoc.SetId)
	if err != nil {
		panic(err)
	}

	ownerId, err := bson.ObjectIDFromHex(helperDoc.OwnerId)
	if err != nil {
		panic(err)
	}

	collection := models.TrackerCollection()

	filter := bson.M{"owner": ownerId}
	update := bson.M{"$pull": bson.M{"remaining_sets": setId}}
	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Item not found dawg"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Removed item dawg"})
}

func ResetTracker(c *gin.Context) {
	_id, err := bson.ObjectIDFromHex(c.Param("param1"))

	if err != nil {
		panic(err)
	}

	filter := bson.M{"owner": _id}
	opts := options.Find().SetProjection(bson.M{
		"_id":   1,
		"owner": 1,
	})

	collection := models.NgSetCollection()

	cursor, err := collection.Find(context.TODO(), filter, opts)

	if err != nil {
		panic(err)
	}

	var results []models.NgSetSchema

	cursor.All(context.TODO(), &results)

	var setIds []bson.ObjectID

	for _, v := range results {
		setIds = append(setIds, v.Id)

		// fmt.Printf("%d %s\n", i, v)
	}

	collection = models.TrackerCollection()

	filter = bson.M{"owner": _id}
	update := bson.D{{"$set", bson.D{{"remaining_sets", setIds}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Tracker reset for " + _id.Hex(),
		// "result":  result,
	})
}
