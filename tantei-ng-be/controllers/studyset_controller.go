package controllers

import (
	"context"
	"net/http"
	"slices"
	"tantei-ng/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func GetNgSets(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var paramOwner = c.Param("owner")
	var ownerObjectID, err = bson.ObjectIDFromHex(paramOwner)

	if err != nil {
		panic(err)
	}

	var collection *mongo.Collection = models.StudysetCollection()

	var opts = options.Find().SetProjection(bson.D{{"items", 0}})
	filter := bson.D{{"owner", ownerObjectID}}
	cursor, err := collection.Find(context.TODO(), filter, opts)

	var results []models.StudysetSchema

	err = cursor.All(context.TODO(), &results)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, results)
}

func GetNgSet(c *gin.Context) {
	var collection *mongo.Collection = models.StudysetCollection()
	var studysetParam = c.Param("studyset")

	var setObjectID, err = bson.ObjectIDFromHex(studysetParam)

	if err != nil {
		panic(err)
	}

	var filter bson.D = bson.D{{"_id", setObjectID}}

	var result models.StudysetSchema
	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, result)
}

func CreateNgSet(c *gin.Context) {
	var collection *mongo.Collection = models.StudysetCollection()

	var doc models.StudysetSchema

	var err = c.BindJSON(&doc)

	var detectedRadicals []string

	if len(doc.Items) > 0 {
		for _, value := range doc.Items {
			for _, rad := range value.Radical {
				if !slices.Contains(detectedRadicals, rad) {
					detectedRadicals = append(detectedRadicals, rad)
				}
			}

		}
	}

	doc.IndexedRadicals = detectedRadicals

	if err != nil {
		panic(err)
	}

	collection.InsertOne(context.TODO(), doc)

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message": "Created new set!",
	})
}

func AddNgSetItems(c *gin.Context) {
	var collection mongo.Collection = *models.StudysetCollection()
	var studysetParam = c.Param("studyset")

	var docs []models.StudywordSchema

	var err = c.BindJSON(&docs)

	if err != nil {
		panic(err)
	}

	studySetObjectID, err := bson.ObjectIDFromHex(studysetParam)

	if err != nil {
		panic(err)
	}

	var detectedRadicals []string
	for _, value := range docs {
		for _, rad := range value.Radical {
			detectedRadicals = append(detectedRadicals, rad)
		}

	}

	// var result models.StudysetSchema

	filter := bson.D{{"_id", studySetObjectID}}
	update := bson.D{{"$push", bson.D{{"items", bson.D{{"$each", docs}}}}}, {"$addToSet", bson.D{{"indexed_radicals", bson.D{{"$each", detectedRadicals}}}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	// c.IndentedJSON(http.StatusOK, result)
	c.IndentedJSON(http.StatusOK, gin.H{"msg": "Successfully added items!"})
}

func RemoveNgSetItems(c *gin.Context) {
	var collection mongo.Collection = *models.StudysetCollection()
	var studysetParam = c.Param("studyset")

	var docs []models.StudywordSchema

	var err = c.BindJSON(&docs)

	if err != nil {
		panic(err)
	}

	studySetObjectID, err := bson.ObjectIDFromHex(studysetParam)

	if err != nil {
		panic(err)
	}

	// var result models.StudysetSchema

	filter := bson.D{{"_id", studySetObjectID}}
	update := bson.D{{"$pull", bson.D{{"items", bson.D{{"$each", docs}}}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	// c.IndentedJSON(http.StatusOK, result)
	c.IndentedJSON(http.StatusOK, gin.H{"msg": "Successfully removed items!"})
}

func TryIndexingRadicalsStudyset(c *gin.Context) {
	studysetParam := c.Param("studyset")
	studysetObjectID, err := bson.ObjectIDFromHex(studysetParam)

	if err != nil {
		panic(err)
	}

	collection := models.StudysetCollection()

	filter := bson.D{{"_id", studysetObjectID}}

	var doc models.StudysetSchema

	err = collection.FindOne(context.TODO(), filter).Decode(&doc)

	if err != nil {
		panic(err)
	}

	var detectedRadicals []string

	for _, val := range doc.Items {
		for _, rad := range val.Radical {
			if !slices.Contains(detectedRadicals, rad) {
				detectedRadicals = append(detectedRadicals, rad)
			}
		}
	}

	update := bson.D{{"$addToSet", bson.D{{"indexed_radicals", bson.D{{"$each", detectedRadicals}}}}}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		panic(err)
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "Radicals are successfully indexed for this studyset!"})
}

func GetCumulativeRadicalsFromOwnedStudyset(c *gin.Context) {
	ownerParam := c.Param("owner")
	ownerObjectID, err := bson.ObjectIDFromHex(ownerParam)

	if err != nil {
		panic(err)
	}

	collection := models.StudysetCollection()

	filter := bson.D{{"owner", ownerObjectID}}
	opts := options.Find().SetProjection(bson.D{{"owner", 0}, {"name", 0}, {"_id", 0}})

	var docs []models.StudysetSchema

	cursor, err := collection.Find(context.TODO(), filter, opts)

	err = cursor.All(context.TODO(), &docs)

	var unique_owned_radicals []string

	for _, doc := range docs {
		for _, radical := range doc.IndexedRadicals {
			if !slices.Contains(unique_owned_radicals, radical) {
				unique_owned_radicals = append(unique_owned_radicals, radical)
			}
		}
	}

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "Success!", "payload": unique_owned_radicals})
}
