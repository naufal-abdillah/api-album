package config

import (
	"database/sql"
	"example/web-service-gin/models"
	"fmt"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SqlQuery() []models.Album {
	db, err := connect()
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
	fmt.Println(len(output))
	return output
}
