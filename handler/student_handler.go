package handler

import (
	"example/web-service-gin/services"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	c.IndentedJSON(services.ServicesGetAlbums())
}
