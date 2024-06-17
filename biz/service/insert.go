package service

import (
	"context"
	"github.com/pkg/errors"
	"io"
	"milvus-demo/biz/rpc"
	"milvus-demo/config"
	"milvus-demo/dal/db"
	"milvus-demo/dal/milvus"
	"milvus-demo/dal/minio"
	pb "milvus-demo/proto"
	"mime/multipart"
)

func (s *PictureService) Insert(ctx context.Context, file *multipart.FileHeader) (image *db.Image, err error) {
	//现将数据上传到minio
	err = minio.UploadData(ctx, file)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert upload data error")
	}

	//保存到mysql
	//确保上传成功后再写url保存到mysql，保持数据一致
	conf := config.Config.Minio
	url := "http://" + conf.Endpoint + "/" + conf.BucketName + "/" + file.Filename
	imageDao := db.NewImageDao(ctx)
	image = &db.Image{
		URL: url,
	}
	err = imageDao.CreateImage(image)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert create image error")
	}
	//保存到milvus
	f, _ := file.Open()
	data, err := io.ReadAll(f)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert read data error")
	}
	//rpc调用获取向量数据
	vector, err := rpc.GetImageVector(ctx, &pb.ImageRequest{Image: data})
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert get vector error")
	}
	err = milvus.InsertVector(ctx, vector, image.Id)
	if err != nil {
		return nil, errors.WithMessage(err, "service.Insert insert vector error")
	}
	return image, nil
}
