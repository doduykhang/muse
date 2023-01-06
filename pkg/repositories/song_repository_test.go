package repositories

import (
	"fmt"
	"testing"
)

var (
	testSongRepository = NewSongRepository()
)

func TestGetSongOfAlbums(t *testing.T) {
	results, err := testSongRepository.GetSongsOfAlbums(1)

	if err != nil {
		t.Errorf("error %v", err)
	}

	fmt.Printf("id 1 %v", results)

	results, err = testSongRepository.GetSongsOfAlbums(2)

	if err != nil {
		t.Errorf("error %v", err)
	}

	fmt.Printf("id 2 %v", results)
}
