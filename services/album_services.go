package services

import (
	"example/web-service-gin/interfaces"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"
	"fmt"
	"regexp"
	"strconv"
)

type AlbumService struct {
}

func (S AlbumService) ServicesGetAlbums() (int, []models.Album) {
	var IRepo interfaces.IAlbumRepo = repositories.AlbumRepo{}
	return IRepo.RepoGetAlbum()
}

func (S AlbumService) ServicesGetAlbumById(param string) (int, []models.Album) {
	var IRepo interfaces.IAlbumRepo = repositories.AlbumRepo{}
	var id int = checkParamId(param)
	return IRepo.RepoGetAlbumById(id)
}
func (S AlbumService) ServicesAddAlbum(input models.Album) {
	var IRepo interfaces.IAlbumRepo = repositories.AlbumRepo{}
	IRepo.RepoAddAlbum(input)
}
func (S AlbumService) ServicesUpdateAlbum(param string, input models.Album) {
	var id int = checkParamId(param)
	var IRepo interfaces.IAlbumRepo = repositories.AlbumRepo{}
	IRepo.RepoUpdateAlbum(id, input)
}

func checkParamId(inputString string) (outputInt int) {
	var regex, _ = regexp.Compile(`[1-9][0-9]*`)
	if regex.MatchString(inputString) {
		outputInt, _ = strconv.Atoi(inputString)

	} else {
		fmt.Print("Parameter Invalid")
	}
	return outputInt
}
