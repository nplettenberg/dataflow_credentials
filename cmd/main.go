package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	mongo, err := CreateMongoConnection()
	if err != nil {
		panic(err)
	}

	collection := mongo.Database("credentials").Collection("secrets")

	defer func() {
		if err = mongo.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/preview", GetCredentialsPreviewHandler(collection))
		v1.GET("/:key", GetSingleCredentialsHandler(collection))
		v1.POST("/", CreateCredentialsHandler(collection))
		v1.DELETE("/:key", DeleteCredentialsHandler(collection))
	}

	log.Fatal(router.Run(":80"))
}
