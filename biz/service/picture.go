package service

import "sync"

var PictureServiceOnce sync.Once

var PictureServiceIns *PictureService

type PictureService struct {
}

func GetPictureService() *PictureService {
	PictureServiceOnce.Do(func() {
		PictureServiceIns = &PictureService{}
	})
	return PictureServiceIns
}
