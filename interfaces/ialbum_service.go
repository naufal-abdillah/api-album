package interfaces

import (
	"example/web-service-gin/models"
)

type IAlbumService interface {
	ServicesGetAlbums() (int, []models.Album)
	ServicesGetAlbumById(Param string) (int, []models.Album)
	ServicesAddAlbum(input models.Album)
	ServicesUpdateAlbum(param string, input models.Album)
}
