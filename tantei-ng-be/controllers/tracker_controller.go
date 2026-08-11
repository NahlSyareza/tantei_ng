package controllers

import (
	"context"
	"fmt"
	"net/http"
	"tantei-ng/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetOwnedSetsAndSave(setIds *[]bson.ObjectID, _id bson.ObjectID) {
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

	for _, v := range results {
		*setIds = append(*setIds, v.Id)

		// fmt.Printf("%d %s\n", i, v)
	}

}

func RefreshCheckTracker(c *gin.Context) {
	_id, err := bson.ObjectIDFromHex(c.Param("param1"))

	if err != nil {
		panic(err)
	}

	filter := bson.D{{"owner", _id}}

	collection := models.TrackerCollection()

	var doc models.TrackerSchema
	err = collection.FindOne(context.TODO(), filter).Decode(&doc)

	if err != nil {
		panic(err)
	}

	currentTime := time.Now()

	elapsedTime := currentTime.Sub(doc.LastTime)

	if elapsedTime >= 24*time.Hour && doc.RefreshType == "Daily" {
		fmt.Printf("Daily reset triggered, blyat")
		var setIds []bson.ObjectID
		GetOwnedSetsAndSave(&setIds, _id)

		collection := models.TrackerCollection()

		filter := bson.M{"owner": _id}
		update := bson.D{{"$set", bson.D{{"remaining_sets", setIds}, {"last_time", time.Now()}}}}

		_, err = collection.UpdateOne(context.TODO(), filter, update)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"refresh_status": "Daily",
			"message":        "Refreshed tracker since it has been one day or more",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Nothing to see here, blyat",
	})
}

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

	// START OWNED SET LOOP
	// filter := bson.M{"owner": _id}
	// opts := options.Find().SetProjection(bson.M{
	// 	"_id":   1,
	// 	"owner": 1,
	// })

	// collection := models.NgSetCollection()

	// cursor, err := collection.Find(context.TODO(), filter, opts)

	// if err != nil {
	// 	panic(err)
	// }

	// var results []models.NgSetSchema

	// cursor.All(context.TODO(), &results)

	// var setIds []bson.ObjectID

	// for _, v := range results {
	// 	setIds = append(setIds, v.Id)

	// 	// fmt.Printf("%d %s\n", i, v)
	// }

	// END OWNED SETS LOOP

	var setIds []bson.ObjectID

	GetOwnedSetsAndSave(&setIds, _id)

	collection := models.TrackerCollection()

	filter := bson.M{"owner": _id}
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
