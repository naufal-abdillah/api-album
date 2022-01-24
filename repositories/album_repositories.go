package repositories

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var err error

func init() {
	if db == nil {
		db, err = config.Connect()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

type AlbumRepo struct {
}

func (R AlbumRepo) RepoGetAlbum() (int, []models.Album) {
	var result []models.Album
	rows, err := db.Query("select * from tb_album")
	if err != nil {
		fmt.Println(err.Error())
		return http.StatusBadRequest, result
	}
	defer rows.Close()

	for rows.Next() {
		var each = models.Album{}
		var err = rows.Scan(&each.ID, &each.Title, &each.Artist, &each.Price)

		if err != nil {
			fmt.Println(err.Error())
			return http.StatusBadRequest, result
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return http.StatusBadRequest, result
	}
	output := make([]models.Album, len(result))
	output = append(output, result...)
	return http.StatusOK, output
}

func (R AlbumRepo) RepoGetAlbumById(Id int) (int, []models.Album) {
	rows, err := db.Query("SELECT * FROM tb_album WHERE ID=?", Id)
	var result []models.Album
	if err != nil {
		fmt.Println(err.Error())
		return http.StatusBadRequest, result
	}
	defer rows.Close()

	for rows.Next() {
		var each = models.Album{}
		var err = rows.Scan(&each.ID, &each.Title, &each.Artist, &each.Price)
		if err != nil {
			fmt.Println(err.Error())
			return http.StatusBadRequest, result
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return http.StatusBadRequest, result

	}
	output := make([]models.Album, len(result))
	output = append(output, result...)

	return http.StatusOK, output
}

func (R AlbumRepo) RepoAddAlbum(input models.Album) {
	entry := `INSERT INTO tb_album (id, title, artist, price) VALUES (?, ?, ?, ?)`
	db.MustExec(entry, input.ID, input.Title, input.Artist, input.Price)

}

func (R AlbumRepo) RepoUpdateAlbum(id int, input models.Album) {
	entry := `UPDATE tb_album SET id=?, title=?, artist=?, price=? WHERE ID=?`
	db.MustExec(entry, input.ID, input.Title, input.Artist, input.Price, id)
}
