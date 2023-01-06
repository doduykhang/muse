package repositories

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
)

type ArtistRepository interface {
	CrudRepository[*models.Aritst, uint]
	SearchArtist(paginate dtos.Paginate, keyword string) (*[]dtos.SeachArtistsResponse, error)
}

type aritstRepository struct {
	CrudRepositoryImpl[*models.Aritst, uint]
}

func (repositoy *aritstRepository) SearchArtist(paginate dtos.Paginate, keyword string) (*[]dtos.SeachArtistsResponse, error) {
	var dtos []dtos.SeachArtistsResponse
	result := db.Raw("call search_artists(?, ?, ?)", paginate.Size, paginate.Page, keyword).Scan(&dtos);
	if result.Error != nil {
		return nil, result.Error
	}
	return &dtos, nil
}

func NewArtistRepository() ArtistRepository {
	return &aritstRepository{}
}
