package controllers

import (
	"github.com/doduykhang/muse/pkg/dtos"
)

type genreController struct {
	ReadController[dtos.GenreDTO, dtos.BaseID]	
} 

func init() {
	songController.service = songService
	songController.CreateUploadController.service = songService
	songController.UpdateController.service = songService
	songController.DeleteController.service = songService
	songController.ReadController.service = songService
}

func NewGenreController() *genreController {
	genreController := &genreController{}
	genreController.ReadController.service = businessServices.GenreService
	return genreController
}

