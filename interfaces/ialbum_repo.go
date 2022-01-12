package interfaces

import (
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

type IAlbumRepo interface {
	RepoGetAlbum() (int, []models.Album)
	RepoGetAlbumById(Id string) (int, []models.Album)
	RepoAddAlbum(c *gin.Context)
	RepoUpdateAlbum(c *gin.Context)
}
