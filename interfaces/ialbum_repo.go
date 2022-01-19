package interfaces

import (
	"example/web-service-gin/models"
)

type IAlbumRepo interface {
	RepoGetAlbum() (int, []models.Album)
	RepoGetAlbumById(Id int) (int, []models.Album)
	RepoAddAlbum(input models.Album)
	RepoUpdateAlbum(id int, input models.Album)
}
