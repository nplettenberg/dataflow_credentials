package main

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func CreateMinioConnection() (*minio.Client, error) {

	minioEndpoint := GetEnv("MINIO_ENDPOINT", "")
	minioKey := GetEnv("MINIO_ACCESS_KEY", "")
	minioSecret := GetEnv("MINIO_ACCESS_SECRET", "")

	fmt.Printf("Connecting to minio:\n Endpoint: %s \n KeyLenght: %d \n SecretLenght: %d \n", minioEndpoint, len(minioKey), len(minioSecret))

	if len(minioEndpoint) == 0 || len(minioKey) == 0 || len(minioSecret) == 0 {
		panic("Missing connection information for MinIO")
	}

	return minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioKey, minioSecret, ""),
		Secure: false,
	})

}
