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

func GetAlbumById(c *gin.Context) {
	var IService interfaces.IAlbumService
	IService = services.Service{}
	// IService.ServicesGetAlbums = services.ServicesGetAlbums
	// c.IndentedJSON(services.ServicesGetAlbums())
	c.IndentedJSON(IService.ServicesGetAlbumById(c))
	// var Status, task = IService.ServicesGetAlbumById(c)
	// c.JSON(Status, gin.H{"data": task})
}
func AddAlbum(c *gin.Context) {
	var IService interfaces.IAlbumService
	IService = services.Service{}
	// IService.ServicesGetAlbums = services.ServicesGetAlbums
	// c.IndentedJSON(services.ServicesGetAlbums())
	// c.IndentedJSON(IService.ServicesAddAlbum(c))
	IService.ServicesAddAlbum(c)

	// var Status, task = IService.ServicesGetAlbumById(c)
	// c.JSON(Status, gin.H{"data": task})
}
func UpdateAlbum(c *gin.Context) {
	var IService interfaces.IAlbumService
	IService = services.Service{}
	// IService.ServicesGetAlbums = services.ServicesGetAlbums
	// c.IndentedJSON(services.ServicesGetAlbums())
	IService.ServicesUpdateAlbum(c)
	// c.IndentedJSON(IService.ServicesUpdateAlbum(c))

	// var Status, task = IService.ServicesGetAlbumById(c)
	// c.JSON(Status, gin.H{"data": task})
}
