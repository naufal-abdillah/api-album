package interfaces

import "example/web-service-gin/models"

type IAlbumService interface {
	ServicesGetAlbums() (int, []models.Album)
}
