package repositories

import (
	"example/web-service-gin/config"
	"example/web-service-gin/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

// func Query(tableName string, condition string) []models.Album {
// 	// var query string
// 	Query := `SELECT * FROM tableName WHERE condition`
// 	fmt.Print(tableName)
// 	rows, err := db.Query(Query)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer rows.Close()

// 	var result []models.Album

// 	for rows.Next() {
// 		var each = models.Album{}
// 		var err = rows.Scan(&each.ID, &each.Title, &each.Artist, &each.Price)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		result = append(result, each)
// 	}

// 	if err = rows.Err(); err != nil {
// 		fmt.Println(err.Error())
// 		// return
// 	}
// 	output := make([]models.Album, len(result))
// 	output = append(output, result...)
// 	return output
// }

// func SqlQuery() []models.Album {
// 	rows, err := db.Query("select * from tb_album")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer rows.Close()

// 	var result []models.Album

// 	for rows.Next() {
// 		var each = models.Album{}
// 		var err = rows.Scan(&each.ID, &each.Title, &each.Artist, &each.Price)

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		result = append(result, each)
// 	}

// 	if err = rows.Err(); err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	output := make([]models.Album, len(result))
// 	output = append(output, result...)
// 	return output
// }

// func SqlQueryById(c *gin.Context) []models.Album {
// 	var Param string = c.Param("id")

// 	rows, err := db.Query("SELECT * FROM tb_album WHERE ID=?", Param)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	defer rows.Close()

// 	var result []models.Album

// 	for rows.Next() {
// 		var each = models.Album{}
// 		var err = rows.Scan(&each.ID, &each.Title, &each.Artist, &each.Price)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		result = append(result, each)
// 	}

// 	// Trying adding entry
// 	// var temp = models.Album{ID: "4", Title: "An Evening With Silk Sonic", Artist: "Silk Sonic", Price: 12.42}
// 	// entry := `INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)`
// 	// db.MustExec(entry, temp.ID, temp.Title, temp.Artist, temp.Price)

// 	if err = rows.Err(); err != nil {
// 		fmt.Println(err.Error())
// 		// return
// 	}
// 	output := make([]models.Album, len(result))
// 	// for _, each := range result {
// 	// 	output = append(output, each)
// 	// }
// 	output = append(output, result...)
// 	fmt.Print(len(output))

// 	return output
// }

// func SqlQueryUpdate(c *gin.Context) []models.Album {
// 	var Param string = c.Param("id")
// 	var Album = SqlQuery()
// 	var input models.Album
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}
// 	// Trying adding entry
// 	// var temp = models.Album{ID: "4", Title: "An Evening With Silk Sonic", Artist: "Silk Sonic", Price: 12.42}
// 	entry := `UPDATE tb_album SET id=?, title=?, artist=?, price=? WHERE ID=?`
// 	db.MustExec(entry, input.ID, input.Title, input.Artist, input.Price, Param)
// 	Album = SqlQuery()
// 	return Album
// }

type Repo struct {
}

// func (R Repo) RepoGetAlbum() (int, []models.Album) {
// 	var Albums = Query("tb_album", "1")
// 	return http.StatusOK, Albums
// }
func (R Repo) RepoGetAlbum() (int, []models.Album) {
	rows, err := db.Query("select * from tb_album")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []models.Album

	for rows.Next() {
		var each = models.Album{}
		var err = rows.Scan(&each.ID, &each.Title, &each.Artist, &each.Price)

		if err != nil {
			fmt.Println(err.Error())
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		// return
	}
	output := make([]models.Album, len(result))
	output = append(output, result...)
	return http.StatusOK, output
}

// func (R Repo) RepoGetAlbumById(c *gin.Context) (int, []models.Album) {
// 	var Album = SqlQueryById(c)
// 	return http.StatusOK, Album
// }

func (R Repo) RepoGetAlbumById(c *gin.Context) (int, []models.Album) {
	var Param string = c.Param("id")

	rows, err := db.Query("SELECT * FROM tb_album WHERE ID=?", Param)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []models.Album

	for rows.Next() {
		var each = models.Album{}
		var err = rows.Scan(&each.ID, &each.Title, &each.Artist, &each.Price)
		if err != nil {
			fmt.Println(err.Error())
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
	// fmt.Print(len(output))

	return http.StatusOK, output
}

func (R Repo) RepoAddAlbum(c *gin.Context) {

	// var Album = SqlQuery()
	var input models.Album
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return http.StatusBadRequest, Album
	}
	// Trying adding entry
	// var temp = models.Album{ID: "4", Title: "An Evening With Silk Sonic", Artist: "Silk Sonic", Price: 12.42}
	entry := `INSERT INTO tb_album (id, title, artist, price) VALUES (?, ?, ?, ?)`
	db.MustExec(entry, input.ID, input.Title, input.Artist, input.Price)
	// update Album variable after query
	// Album = SqlQuery()
	// fmt.Print(input.ID)
	// fmt.Print("Printing Adding")
	// return http.StatusOK, Album
}

// func (R Repo) RepoUpdateAlbum(c *gin.Context) (int, []models.Album) {
// 	SqlQueryUpdate(c)
// 	var Album = SqlQuery()
// 	// fmt.Print(input.ID)
// 	// fmt.Print("Printing Adding")
// 	return http.StatusOK, Album
// }

func (R Repo) RepoUpdateAlbum(c *gin.Context) {
	var Param string = c.Param("id")
	// var Album = SqlQuery()
	var input models.Album
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	// Trying adding entry
	// var temp = models.Album{ID: "4", Title: "An Evening With Silk Sonic", Artist: "Silk Sonic", Price: 12.42}
	entry := `UPDATE tb_album SET id=?, title=?, artist=?, price=? WHERE ID=?`
	db.MustExec(entry, input.ID, input.Title, input.Artist, input.Price, Param)
	// Album = SqlQuery()
	// return Album
}
