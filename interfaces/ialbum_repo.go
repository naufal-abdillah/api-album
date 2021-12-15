package interfaces

import "example/web-service-gin/models"

type IAlbumRepo interface {
	RepoGetAlbum() (int, []models.Album)
}
