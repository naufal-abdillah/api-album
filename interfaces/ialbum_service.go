package interfaces

import (
	"example/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

type IAlbumService interface {
	ServicesGetAlbums() (int, []models.Album)
	ServicesGetAlbumById(c *gin.Context) (int, []models.Album)
	ServicesAddAlbum(c *gin.Context) (int, []models.Album)
	ServicesUpdateAlbum(c *gin.Context) (int, []models.Album)
}
