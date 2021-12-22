package repositories

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"fmt"
	"net/http"
)

func SqlQuery() []models.Album {
	db, err := config.Connect()
	if err != nil {
		fmt.Println(err.Error())
		// return
	}
	defer db.Close()

	rows, err := db.Query("select * from tb_album")
	if err != nil {
		fmt.Println(err.Error())
		// return
	}
	defer rows.Close()

	var result []models.Album

	for rows.Next() {
		var each = models.Album{}
		var err = rows.Scan(&each.ID, &each.Title, &each.Artist, &each.Price)

		if err != nil {
			fmt.Println(err.Error())
			// return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		// return
	}
	output := make([]models.Album, len(result))
	for _, each := range result {
		output = append(output, each)
	}
	return output
}

type Repo struct {
}

func (R Repo) RepoGetAlbum() (int, []models.Album) {
	var Albums = SqlQuery()
	return http.StatusOK, Albums
}
