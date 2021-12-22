package handler

import (
	"example/web-service-gin/interfaces"
	"example/web-service-gin/services"

	// "example/web-service-gin/services"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var IService interfaces.IAlbumService
	IService = services.Service{}
	// IService.ServicesGetAlbums = services.ServicesGetAlbums
	// c.IndentedJSON(services.ServicesGetAlbums())
	c.IndentedJSON(IService.ServicesGetAlbums())
}