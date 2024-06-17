package rpc

import pb "milvus-demo/proto"

var (
	convertClient pb.ClipServiceClient
)

func Init() {
	InitConvertRPC()
}
