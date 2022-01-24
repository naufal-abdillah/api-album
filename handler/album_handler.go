package handler

import (
	"example/web-service-gin/interfaces"
	"example/web-service-gin/models"
	"example/web-service-gin/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerGetAlbums(c *gin.Context) {
	var IService interfaces.IAlbumService = services.AlbumService{}
	c.IndentedJSON(IService.ServicesGetAlbums())
}

func HandlerGetAlbumById(c *gin.Context) {
	var Param string = c.Param("id")
	var IService interfaces.IAlbumService = services.AlbumService{}
	c.IndentedJSON(IService.ServicesGetAlbumById(Param))
}
func HandlerAddAlbum(c *gin.Context) {
	var input models.Album
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var IService interfaces.IAlbumService = services.AlbumService{}
	IService.ServicesAddAlbum(input)
}
func HandlerUpdateAlbum(c *gin.Context) {
	var param string = c.Param("id")
	var input models.Album
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var IService interfaces.IAlbumService = services.AlbumService{}
	IService.ServicesUpdateAlbum(param, input)
}
