package main

// import "example/web-service-gin/routes"
import (
	"example/web-service-gin/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.Run()
}
