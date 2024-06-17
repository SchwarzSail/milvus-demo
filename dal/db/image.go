package db

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Image struct {
	Id  int64  `gorm:"primary_key;AUTO_INCREMENT"`
	URL string `gorm:"column:url"`
}

type ImageDao struct {
	*gorm.DB
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := DB
	return db.WithContext(ctx)
}

func NewImageDao(ctx context.Context) *ImageDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &ImageDao{NewDBClient(ctx)}
}

func (dao *ImageDao) CreateImage(image *Image) (err error) {
	err = dao.Model(&Image{}).Create(image).Error
	if err != nil {
		return errors.Wrap(err, "failed to create image")
	}
	return nil
}

func (dao *ImageDao) GetUrlByID(id int64) (image *Image, err error) {
	err = dao.Model(&Image{}).Where("id = ?", id).First(&image).Error
	if err != nil {
		return nil, errors.Wrap(err, "failed to get image")
	}
	return image, nil
}
