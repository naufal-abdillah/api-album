package services

import (
	"example/web-service-gin/interfaces"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"

	"github.com/gin-gonic/gin"
)

type Service struct {
}

func (S Service) ServicesGetAlbums() (int, []models.Album) {
	var IRepo interfaces.IAlbumRepo
	IRepo = repositories.Repo{}
	return IRepo.RepoGetAlbum()
}

func (S Service) ServicesGetAlbumById(c *gin.Context) (int, []models.Album) {
	var IRepo interfaces.IAlbumRepo
	IRepo = repositories.Repo{}
	// handle kalau nol/huruf

	return IRepo.RepoGetAlbumById(c)
}
func (S Service) ServicesAddAlbum(c *gin.Context) {
	var IRepo interfaces.IAlbumRepo
	IRepo = repositories.Repo{}
	IRepo.RepoAddAlbum(c)
}
func (S Service) ServicesUpdateAlbum(c *gin.Context) {
	var IRepo interfaces.IAlbumRepo
	IRepo = repositories.Repo{}
	IRepo.RepoUpdateAlbum(c)
}
