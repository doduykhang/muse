package repositories

import (
	"github.com/doduykhang/muse/pkg/models"
)

type GenreRepository interface {
	CrudRepository[models.Genre, models.BaseID]
}

type genreRepository struct {
	CrudRepositoryImpl[models.Genre, models.BaseID]
}

func NewGenereRepository() GenreRepository {
	return &genreRepository{}
}
