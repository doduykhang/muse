package repositories

import (
	"fmt"
	"testing"
)

var (
	testAlbumRepository = NewAlbumRepository()
)

func TestGetAlbumsOfArtist(t *testing.T) {
	result, err := testAlbumRepository.GetAlbumsOfArtist(1)	
	
	if err != nil {
		t.Errorf("error %v", err)
	}

	fmt.Println(result)

	result, err = testAlbumRepository.GetAlbumsOfArtist(2)	
	
	if err != nil {
		t.Errorf("error %v", err)
	}

	fmt.Println(result)
}

