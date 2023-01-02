package service

import (
	"os"

	"github.com/doduykhang/muse/pkg/dtos"
)

type FileService interface {
	UploadFiles([]*dtos.FileDTO) (map[string][]string, error)
}

func NewFileService() FileService {
	return &localFileService{}
}

type localFileService struct{}

func (service *localFileService) UploadFiles(fileDtos []*dtos.FileDTO) (map[string][]string, error) {
	m := make(map[string][]string)
	for _, file := range fileDtos {
		err := os.WriteFile("assets"+file.Path+file.Name, file.Bytes, 0644)

		if err != nil {
			return nil, err
		}
		m[file.Field] = []string{"assets" + file.Path + file.Name}
	}
	return m, nil
}
