package routes

import (
	"example/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumById)
	router.Run("localhost:8000")
}
