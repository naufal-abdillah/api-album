package routes

import (
	"example/web-service-gin/handler"

	"github.com/gin-gonic/gin"
)

func Routes() {
	router := gin.Default()
	router.GET("/albums", handler.GetAlbums)
	router.Run("localhost:8000")
}
