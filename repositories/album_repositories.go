package repositories

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"fmt"
	"net/http"
)

type Repo struct {
}

func (R Repo) RepoGetAlbum() (int, []models.Album) {
	var Albums = config.SqlQuery()
	fmt.Println("---------Print Repo----------")
	return http.StatusOK, Albums
}
