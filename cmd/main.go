package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	client, err := CreateMinioConnection()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/preview", GetCredentialsPreviewHandler(client))
		v1.GET("/:key", GetSingleCredentialsHandler(client))
		v1.POST("/", CreateCredentialsHandler(client))
		v1.DELETE("/:key", DeleteCredentialsHandler(client))
	}

	log.Fatal(router.Run(":80"))
}
