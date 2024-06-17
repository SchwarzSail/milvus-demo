package dal

import (
	"milvus-demo/dal/db"
	"milvus-demo/dal/milvus"
	"milvus-demo/dal/minio"
)

func Init() {
	milvus.Init()
	minio.Init()
	db.InitMySQL()
}
