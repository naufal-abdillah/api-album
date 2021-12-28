package main

// import "example/web-service-gin/routes"
import (
	"example/web-service-gin/app"

	_ "github.com/go-sql-driver/mysql"
)

// var schema = `
// CREATE TABLE person (
//     first_name text,
//     last_name text,
//     email text
// )
// `

// type Person struct {
// 	FirstName string `db:"first_name"`
// 	LastName  string `db:"last_name"`
// 	Email     string
// }

// type Place struct {
// 	Country string
// 	City    sql.NullString
// 	TelCode int
// }

func main() {
	// db, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	// return
	// }

	// tx := db.MustBegin()

	// result, _ := db.Query("SHOW TABLES")
	// var table string
	// for result.Next() {
	// 	result.Scan(&table)
	// 	fmt.Println(table)
	// }
	app.Run()

}
