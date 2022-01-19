package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func Connect() (*sqlx.DB, error) {
	// db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	// if err != nil {
	// 	return nil, err
	// }

	// return db, nil
	db, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		fmt.Println(err.Error())
		// return
	}
	return db, nil
}
