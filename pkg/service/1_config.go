package service

import (
	"github.com/doduykhang/muse/pkg/repositories"
	"github.com/dranikpg/dto-mapper"
	"github.com/gorilla/schema"
)

var (
	mapper         dto.Mapper
	decoder        = schema.NewDecoder()
	appRepositores = repositories.GetAppRepositories()
)

func init() {
	mapper = dto.Mapper{}
}
