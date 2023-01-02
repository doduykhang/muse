package repositories

import (
	"github.com/doduykhang/muse/pkg/models"
)

type ArtistRepository interface {
	CrudRepository[*models.Aritst, uint]
}

type aritstRepository struct {
	CrudRepositoryImpl[*models.Aritst, uint]
}

func NewArtistRepository() ArtistRepository {
	return &aritstRepository{}
}
