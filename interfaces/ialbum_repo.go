package interfaces

import (
	"example/web-service-gin/models"
)

type IAlbumRepo interface {
	RepoGetAlbum() (int, []models.Album)
	RepoGetAlbumById(Id string) (int, []models.Album)
	RepoAddAlbum(input models.Album)
	RepoUpdateAlbum(id string, input models.Album)
}
