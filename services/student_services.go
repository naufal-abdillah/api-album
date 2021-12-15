package services

import (
	"example/web-service-gin/interfaces"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"
)

type Service struct {
}

func (S Service) ServicesGetAlbums() (int, []models.Album) {
	var IRepo interfaces.IAlbumRepo
	IRepo = repositories.Repo{}
	// return repositories.RepoGetAlbum()
	return IRepo.RepoGetAlbum()
}
