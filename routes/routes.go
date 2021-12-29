package routes

import (
	"example/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.GET("/albums/:id", handler.GetAlbumById)
	router.POST("/albums", handler.AddAlbum)
	// router.POST("/albums/:id", handler.UpdateAlbum)
	router.Run("localhost:8000")
}
