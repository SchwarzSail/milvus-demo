package minio

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"milvus-demo/config"
)

type Client struct {
	client *minio.Client
}

var MinioClient *Client

func Init() {
	conf := config.Config
	c, err := minio.New(conf.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4("minioadmin", "minioadmin", ""),
		Secure: false,
	})
	if err != nil {
		panic(err)
	}
	client := &Client{
		client: c,
	}
	MinioClient = client
}
