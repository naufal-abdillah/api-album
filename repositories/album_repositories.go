package repositories

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SqlQuery() []models.Album {
	db, err := config.Connect()
	// db, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		fmt.Println(err.Error())
		// return
	}
	defer db.Close()

	// Albums := []models.Album{}
	// db.Select(&Albums, "select * from tb_album")

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
	// for _, each := range result {
	// 	output = append(output, each)
	// }
	output = append(output, result...)
	return output
}

func SqlQueryById(c *gin.Context) []models.Album {
	var Param string
	Param = c.Param("id")

	db, err := config.Connect()
	// db, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		fmt.Println(err.Error())
		// return
	}
	defer db.Close()

	// Albums := []models.Album{}
	// db.Select(&Albums, "select * from tb_album")

	rows, err := db.Query("SELECT * FROM tb_album WHERE ID=?", Param)
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

	// Trying adding entry
	// var temp = models.Album{ID: "4", Title: "An Evening With Silk Sonic", Artist: "Silk Sonic", Price: 12.42}
	// entry := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`
	// db.MustExec(entry, temp.ID, temp.Title, temp.Artist, temp.Price)

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		// return
	}
	output := make([]models.Album, len(result))
	// for _, each := range result {
	// 	output = append(output, each)
	// }
	output = append(output, result...)
	fmt.Print(len(output))

	return output
}

type Repo struct {
}

func (R Repo) RepoGetAlbum() (int, []models.Album) {
	var Albums = SqlQuery()
	return http.StatusOK, Albums
}

func (R Repo) RepoGetAlbumById(c *gin.Context) (int, []models.Album) {
	var Album = SqlQueryById(c)
	return http.StatusOK, Album
}
