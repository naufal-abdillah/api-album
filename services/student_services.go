package services

import (
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"
)

func ServicesGetAlbums() (int, []models.Album) {
	return repositories.RepoGetAlbum()
}
