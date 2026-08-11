package controllers

import (
	"context"
	"errors"
	"net/http"
	"tantei-ng/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func RegisterAccount(c *gin.Context) {
	var collection *mongo.Collection = models.AccountCollection()

	var doc models.AccountSchema

	var err = c.BindJSON(&doc)

	// doc.Id = bson.NewObjectID()
	// doc.CreatedAt = bson.NewObjectID().Timestamp()
	// doc.UpdatedAt = bson.NewObjectID().Timestamp()

	if err != nil {
		panic(err)
	}

	collection.InsertOne(context.TODO(), doc)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Successfully registered new account",
	})
}

type accountLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginAccount(c *gin.Context) {
	var loginDoc accountLogin

	var err = c.BindJSON(&loginDoc)

	if err != nil {
		panic(err)
	}

	var doc models.AccountSchema

	var collection = models.AccountCollection()

	var filter = bson.M{"email": loginDoc.Email}

	err = collection.FindOne(context.TODO(), filter).Decode(&doc)

	if loginDoc.Password != doc.Password {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "Wrong password, blyat",
		})

		return
	}

	OwnedSetsAccount(doc.Id.Hex())

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Correct password, blin",
	})
}

// Integrated into the login func already
// func OwnedSetsByAccount(c* gin.Context){
func OwnedSetsAccount(oid_par string) {
	// Try retrieve ALL created sets from an owner
	// oid is the owner's ObjectID
	// oid, err := bson.ObjectIDFromHex(c.Param("oid"))
	oid, err := bson.ObjectIDFromHex(oid_par)

	if err != nil {
		panic(err)
	}

	filter := bson.M{"owner": oid}
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

	err = cursor.All(context.TODO(), &results)

	if err != nil {
		panic(err)
	}

	if len(results) == 0 {
		// c.JSON(http.StatusOK, gin.H{"message": "This account either doesn't exist or doesn't have any sets owned"})
		return
	}

	var setIds []bson.ObjectID

	for _, v := range results {
		setIds = append(setIds, v.Id)
	}

	// Check for tracker
	trackerCollection := models.TrackerCollection()
	var trackerDoc models.TrackerSchema

	filter = bson.M{"owner": oid}
	err = trackerCollection.FindOne(context.TODO(), filter).Decode(&trackerDoc)

	if errors.Is(err, mongo.ErrNoDocuments) {
		trackerCollection.InsertOne(context.TODO(), models.TrackerSchema{
			Owner:         oid,
			LastTime:      time.Now(),
			RemainingSets: setIds,
		})

		// fmt.Printf("Created new entry\n")
	} else {
		// fmt.Printf("Not creating new entry\n")
	}

	// c.IndentedJSON(http.StatusOK, results)
}

// Remove a recorded set from the tracker
