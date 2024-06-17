package pack

import (
	"milvus-demo/biz/model/image"
	"milvus-demo/dal/db"
)

func BuildImage(i *db.Image) *image.Image {
	return &image.Image{
		ID:  i.Id,
		URL: i.URL,
	}
}

func BuildImageList(images []*db.Image) []*image.Image {
	resp := make([]*image.Image, len(images))
	for i, img := range images {
		resp[i] = BuildImage(img)
	}
	return resp
}
