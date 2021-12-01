package repositories

import (
	"example/web-service-gin/models"
	"net/http"
)

func RepoGetAlbum() (int, []models.Album) {
	var Albums = []models.Album{
		{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
		{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
		{ID: "3", Title: "Walking Slow", Artist: "Mac Ayres", Price: 14.50},
	}
	return http.StatusOK, Albums
}
