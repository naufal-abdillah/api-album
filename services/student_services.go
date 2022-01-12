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
	return IRepo.RepoGetAlbum()
}

func (S Service) ServicesGetAlbumById(Param string) (int, []models.Album) {
	var IRepo interfaces.IAlbumRepo
	IRepo = repositories.Repo{}
	// handle kalau nol/huruf
	var Id string = Param
	return IRepo.RepoGetAlbumById(Id)
}
func (S Service) ServicesAddAlbum(input models.Album) {
	var IRepo interfaces.IAlbumRepo
	IRepo = repositories.Repo{}
	IRepo.RepoAddAlbum(input)
}
func (S Service) ServicesUpdateAlbum(param string, input models.Album) {
	var id string = param
	//check param
	var IRepo interfaces.IAlbumRepo
	IRepo = repositories.Repo{}
	IRepo.RepoUpdateAlbum(id, input)
}
