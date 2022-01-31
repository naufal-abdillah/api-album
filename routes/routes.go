package routes

import (
	"example/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/albums", handler.HandlerGetAlbums)
	router.GET("/albums/:id", handler.HandlerGetAlbumById)
	router.POST("/albums", handler.HandlerAddAlbum)
	router.POST("/albums/:id", handler.HandlerUpdateAlbum)
	// Auth
	router.POST("/user/register", handler.HandlerRegisterUser)
	router.Run("localhost:8000")
}
