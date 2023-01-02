package main

import (
	"fmt"
	"time"

	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
	"github.com/doduykhang/muse/pkg/repositories"
	"github.com/doduykhang/muse/pkg/service"
)

func main() {
	genreRepository := repositories.NewGenereRepository()
	genreRepository.Update(
		models.BaseID{
			ID: 1,
		},
		&models.Genre{
			Name: "genre update",
		},
	)
	//Create()
	//Update()
	//Delete()
	//FindOne()
	//FindAll()
}

func Create() {
	songService := service.NewSongService()
	dto, err := songService.Create(&dtos.CreateSong{
		Title:       "song title",
		Duration:    100,
		ReleaseDate: time.Now(),
		Image:       "image",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("dto %v", dto.ID)
}

func Update() {
	songService := service.NewSongService()
	dto, err := songService.Update(dtos.BaseID{
		ID: 1,
	}, &dtos.UpdateSong{
		Title:       "song title update",
		Duration:    100,
		ReleaseDate: time.Now(),
		Image:       "image",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("dto %v", dto)
}

func Delete() {
	songService := service.NewSongService()
	err := songService.Delete(dtos.BaseID{
		ID: 1,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("deleted")
}

func FindOne() {
	songService := service.NewSongService()
	dto, err := songService.FindOne(dtos.BaseID{
		ID: 1,
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("dto %v", dto)
}

func FindAll() {
	songService := service.NewSongService()
	dto, err := songService.FindAll()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("dto %v", dto)
}
