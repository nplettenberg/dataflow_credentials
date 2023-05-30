package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetCredentialsPreviewHandler(collection *mongo.Collection) func(c *gin.Context) {
	return func(c *gin.Context) {
		cursor, err := collection.Find(context.TODO(), bson.M{})
		if err != nil {
			log.Fatal(err)
		}

		var credentials []CredentialsPreview
		if err = cursor.All(context.TODO(), &credentials); err != nil {
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, credentials)
	}
}

func GetSingleCredentialsHandler(collection *mongo.Collection) func(c *gin.Context) {
	return func(c *gin.Context) {
		key := c.Param("key")

		objectId, err := primitive.ObjectIDFromHex(key)
		if err != nil {
			log.Println("Invalid id")
			c.JSON(http.StatusNotFound, gin.H{
				"error":       err,
				"requestedID": key,
			})
		}

		var credentials Credentials

		if err := collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&credentials); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			fmt.Println(err)
			return
		}

		c.JSON(http.StatusOK, credentials)
	}
}

func CreateCredentialsHandler(collection *mongo.Collection) func(c *gin.Context) {

	return func(c *gin.Context) {
		var credentials Credentials

		err := c.ShouldBind(&credentials)

		if err == nil {

			result, err := collection.InsertOne(context.Background(), credentials)
			if err != nil {
				log.Fatal(err)
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
			}

			c.JSON(http.StatusOK, gin.H{
				"id": result.InsertedID,
			})

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
}

func DeleteCredentialsHandler(collection *mongo.Collection) func(c *gin.Context) {
	return func(c *gin.Context) {
		key := c.Param("key")

		objectId, err := primitive.ObjectIDFromHex(key)

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		result, err := collection.DeleteOne(context.Background(), bson.M{"_id": objectId})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"deletedCount": result.DeletedCount,
		})
	}
}
