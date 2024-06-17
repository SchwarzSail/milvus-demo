package rpc

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "milvus-demo/proto"
)

func InitConvertRPC() {
	conn, err := grpc.Dial("0.0.0.0:50051", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	convertClient = pb.NewClipServiceClient(conn)
}

func GetImageVector(ctx context.Context, req *pb.ImageRequest) (vector []float32, err error) {
	resp, err := convertClient.GetImageVector(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "rpc.GetImageVector failed")
	}
	return resp.Vector, nil
}
func GetTextVector(ctx context.Context, req *pb.TextRequest) (vector []float32, err error) {
	resp, err := convertClient.GetTextVector(ctx, req)
	if err != nil {
		return nil, errors.WithMessage(err, "rpc.GetTextVector failed")
	}
	return resp.Vector, nil
}
