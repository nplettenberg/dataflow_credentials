package main

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
)

func GetCredentialsPreviewHandler(client *minio.Client) func(c *gin.Context) {
	return func(c *gin.Context) {

		var credentials []CredentialsPreview
		for object := range client.ListObjects(context.Background(), "secrets", minio.ListObjectsOptions{}) {

			key := object.Key

			reader, err := client.GetObject(context.Background(), "secrets", key, minio.GetObjectOptions{})

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				fmt.Println(err)
				return
			}

			objectContent, err := ReaderToString(reader)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				fmt.Println(err)
				return
			}

			var credential CredentialsPreview

			if err := json.Unmarshal([]byte(objectContent), &credential); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				fmt.Println(err)
				return
			}

			credentials = append(credentials, credential)
		}

		c.JSON(http.StatusOK, credentials)
	}
}

func ReaderToString(reader io.Reader) (string, error) {

	result, err := io.ReadAll(reader)

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func GetSingleCredentialsHandler(client *minio.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		key := c.Param("key")

		reader, err := client.GetObject(context.Background(), "secrets", fmt.Sprintf("%s.json", key), minio.GetObjectOptions{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			fmt.Println(err)
			return
		}

		defer reader.Close()

		objectContent, err := ReaderToString(reader)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			fmt.Println(err)
			return
		}

		credentials := Credentials{}
		json.Unmarshal([]byte(objectContent), &credentials)

		c.JSON(http.StatusOK, credentials)
	}
}

func CreateCredentialsHandler(client *minio.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		var credentials Credentials

		err := c.ShouldBind(&credentials)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			fmt.Println(err)
			return
		}

		plainHash := fmt.Sprintf("%s.%s", credentials.Name, credentials.Secret)
		byteHash := md5.Sum([]byte(plainHash))
		objectHash := fmt.Sprintf("%x", byteHash)

		credentials.ID = string(objectHash)

		objectContent, err := json.Marshal(credentials)

		if err == nil {
			_, err := client.PutObject(
				context.Background(),
				"secrets",
				fmt.Sprintf("%s.json", objectHash),
				strings.NewReader(string(objectContent)),
				int64(len(objectContent)),
				minio.PutObjectOptions{},
			)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err,
				})
				fmt.Println(err)
				return
			}

			c.JSON(
				http.StatusOK,
				gin.H{
					"createdId": objectHash,
				},
			)

			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
}

func DeleteCredentialsHandler(client *minio.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		key := c.Param("key")

		err := client.RemoveObject(
			context.Background(),
			"secrets",
			fmt.Sprintf("%s.json", key),
			minio.RemoveObjectOptions{},
		)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"deletedId": key,
		})
	}
}
