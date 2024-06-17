package service

import (
	"context"
	"github.com/pkg/errors"
	"milvus-demo/biz/model/image"
	"milvus-demo/biz/rpc"
	"milvus-demo/dal/db"
	"milvus-demo/dal/milvus"
	pb "milvus-demo/proto"
)

func (s *PictureService) SearchByText(ctx context.Context, text string) (images []*image.Image, err error) {
	var ids []int64
	vector, err := rpc.GetTextVector(ctx, &pb.TextRequest{Text: text})
	if err != nil {
		return nil, errors.WithMessage(err, "GetTextVector failed")
	}
	ids, err = milvus.Search(ctx, vector)
	if err != nil {
		return nil, errors.WithMessage(err, "Search failed")
	}
	imageDao := db.NewImageDao(ctx)
	images = make([]*image.Image, 0, len(ids))
	for _, id := range ids {
		picture, err := imageDao.GetUrlByID(id)
		if err != nil {
			return nil, errors.WithMessage(err, "GetUrlByID failed")
		}
		temp := &image.Image{
			ID:  picture.Id,
			URL: picture.URL,
		}
		images = append(images, temp)
	}

	return images, nil
}

func (s *PictureService) SearchByImage(ctx context.Context, data []byte) (images []*image.Image, err error) {
	var ids []int64
	vector, err := rpc.GetImageVector(ctx, &pb.ImageRequest{Image: data})
	if err != nil {
		return nil, errors.WithMessage(err, "GetTextVector failed")
	}
	ids, err = milvus.Search(ctx, vector)
	if err != nil {
		return nil, errors.WithMessage(err, "Search failed")
	}
	imageDao := db.NewImageDao(ctx)
	images = make([]*image.Image, 0, len(ids))
	for _, id := range ids {
		picture, err := imageDao.GetUrlByID(id)
		if err != nil {
			return nil, errors.WithMessage(err, "GetUrlByID failed")
		}
		temp := &image.Image{
			ID:  picture.Id,
			URL: picture.URL,
		}
		images = append(images, temp)
	}
	return images, nil
}
