package main

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func CreateMinioConnection() (*minio.Client, error) {

	minioEndpoint := GetEnv("MINIO_ENDPOINT", "localhost:9001")
	minioKey := GetEnv("MINIO_ACCESS_KEY", "0t7nB2RIqguQXTnPwRTY")
	minioSecret := GetEnv("MINIO_ACCESS_SECRET", "qe1AD0ahdB1eDwgL9uCA3pxjChC4E5Shfpmx2og8")

	return minio.New(minioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioKey, minioSecret, ""),
		Secure: false,
	})

}
