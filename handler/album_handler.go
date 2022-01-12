package handler

import (
	"example/web-service-gin/interfaces"
	"example/web-service-gin/services"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var IService interfaces.IAlbumService
	IService = services.Service{}
	c.IndentedJSON(IService.ServicesGetAlbums())
}

func GetAlbumById(c *gin.Context) {
	var IService interfaces.IAlbumService
	IService = services.Service{}
	c.IndentedJSON(IService.ServicesGetAlbumById(c))
}
func AddAlbum(c *gin.Context) {
	var IService interfaces.IAlbumService
	IService = services.Service{}
	IService.ServicesAddAlbum(c)
}
func UpdateAlbum(c *gin.Context) {
	var IService interfaces.IAlbumService
	IService = services.Service{}
	IService.ServicesUpdateAlbum(c)
}
