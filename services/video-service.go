package services

import "practice/entities"

type VideoService interface {
	FindAll() []entities.Video
	Save(video *entities.Video) entities.Video
}

func New() VideoService {
	return &VideoServiceImpl{
		videos: []entities.Video{},
	}
}

type VideoServiceImpl struct {
	videos []entities.Video
}

func (service *VideoServiceImpl) FindAll() []entities.Video {
	return service.videos
}

func (service *VideoServiceImpl) Save(video *entities.Video) entities.Video {
	service.videos = append(service.videos, *video)
	return service.videos[len(service.videos)-1]
}
