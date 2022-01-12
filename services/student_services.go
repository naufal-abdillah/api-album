package services

import (
	"example/web-service-gin/interfaces"
	"example/web-service-gin/models"
	"example/web-service-gin/repositories"
	"regexp"
	"strconv"
)

type Service struct {
}

func (S Service) ServicesGetAlbums() (int, []models.Album) {
	var IRepo interfaces.IAlbumRepo = repositories.Repo{}
	return IRepo.RepoGetAlbum()
}

func (S Service) ServicesGetAlbumById(param string) (int, []models.Album) {
	var IRepo interfaces.IAlbumRepo = repositories.Repo{}
	var id int = checkParamId(param)
	// handle kalau nol/huruf
	return IRepo.RepoGetAlbumById(id)
}
func (S Service) ServicesAddAlbum(input models.Album) {
	var IRepo interfaces.IAlbumRepo = repositories.Repo{}
	IRepo.RepoAddAlbum(input)
}
func (S Service) ServicesUpdateAlbum(param string, input models.Album) {
	var id int = checkParamId(param)
	//check param
	var IRepo interfaces.IAlbumRepo = repositories.Repo{}
	IRepo.RepoUpdateAlbum(id, input)
}

func checkParamId(inputString string) (outputInt int) {
	var regex, _ = regexp.Compile(`[0-9]*`)
	if inputString != "0" && regex.MatchString(inputString) {
		outputInt, _ = strconv.Atoi(inputString)

	}
	return outputInt
}
