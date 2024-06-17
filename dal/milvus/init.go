package milvus

import (
	"context"
	"github.com/milvus-io/milvus-sdk-go/v2/client"
	"milvus-demo/config"
)

type Client struct {
	milvus client.Client
}

var MClient Client

func Init() {
	conf := config.Config
	ctx := context.Background()
	client, err := client.NewClient(ctx, client.Config{
		Address: conf.Milvus.Endpoint,
	})
	if err != nil {
		panic(err)
	}
	m := Client{
		milvus: client,
	}

	MClient = m

	//顺便建立Collection和Index
	ok, err := CreateCollection(ctx)
	if err != nil {
		panic(err)
	}
	if !ok {
		err = CreateIndex(ctx)
		if err != nil {
			panic(err)
		}
	}
}
