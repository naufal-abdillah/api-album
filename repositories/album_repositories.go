package repositories

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"net/http"
)

type Repo struct {
}

func (R Repo) RepoGetAlbum() (int, []models.Album) {
	var Albums = config.SqlQuery()
	return http.StatusOK, Albums
}
