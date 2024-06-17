package minio

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"milvus-demo/config"
	"mime/multipart"
)

func UploadData(ctx context.Context, file *multipart.FileHeader) (err error) {
	conf := config.Config
	client := MinioClient.client
	objectName := file.Filename
	data, err := file.Open()
	if err != nil {
		return errors.WithMessage(err, "open file error")
	}
	_, err = client.PutObject(ctx, conf.Minio.BucketName, objectName, data, file.Size, minio.PutObjectOptions{
		ContentType: "image/jpeg",
	})
	if err != nil {
		return errors.WithMessage(err, "upload file error")
	}
	return nil
}
