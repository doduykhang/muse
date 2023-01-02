package service

import (
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/repositories"
)

type CreateService[CreateRequest any, DTO any] interface {
	Create(request *CreateRequest) (*DTO, error)
}

type CreateUploadService[CreateRequest any, DTO any] interface {
	CreateUpload(request *CreateRequest, files []*dtos.FileDTO) (*DTO, error)
}

type UpdateService[UpdateRequest any, DTO any, ID any] interface {
	Update(id ID, request *UpdateRequest) (*DTO, error)
}

type DeleteService[ID any] interface {
	Delete(id ID) error
}

type ReadService[DTO any, ID any] interface {
	FindOne(id ID) (*DTO, error)
	FindAll() ([]*DTO, error)
}

type CreateServiceImpl[Model any, ID any, CreateRequest any, DTO any] struct {
	repository repositories.CrudRepository[Model, ID]
}

type CreateUploadServiceImpl[Model any, ID any, CreateRequest any, DTO any] struct {
	repository  repositories.CrudRepository[Model, ID]
	fileService FileService
}

type UpdateSeviceImpl[Model any, ID any, UpdateRequest any, IDRequest any, DTO any] struct {
	repository repositories.CrudRepository[Model, ID]
}

type DeleteServiceImpl[Model any, ID any, IDRequest any] struct {
	repository repositories.CrudRepository[Model, ID]
}

type ReadServiceImpl[Model any, ID any, IDRequest any, DTO any] struct {
	repository repositories.CrudRepository[Model, ID]
}

func (service *CreateServiceImpl[Model, ID, CreateRequest, DTO]) Create(request *CreateRequest) (*DTO, error) {
	var model Model
	mapper.Map(&model, request)

	createdModel, error := service.repository.Create(&model)
	if error != nil {
		return nil, error
	}

	var dto DTO
	mapper.Map(&dto, createdModel)

	return &dto, nil
}

func (service *CreateUploadServiceImpl[Model, ID, CreateRequest, DTO]) CreateUpload(request *CreateRequest, files []*dtos.FileDTO) (*DTO, error) {
	fileUrls, err := service.fileService.UploadFiles(files)
	if err != nil {
		return nil, err
	}

	var model Model
	mapper.Map(&model, request)
	decoder.Decode(&model, fileUrls)

	createdModel, err := service.repository.Create(&model)
	if err != nil {
		return nil, err
	}

	var dto DTO
	mapper.Map(&dto, createdModel)

	return &dto, nil
}

func (service *UpdateSeviceImpl[Model, ID, UpdateRequest, IDRequest, DTO]) Update(idRequest IDRequest, request *UpdateRequest) (*DTO, error) {
	var model Model
	mapper.Map(&model, request)

	var id ID
	mapper.Map(&id, &idRequest)

	updatedModel, error := service.repository.Update(id, &model)
	if error != nil {
		return nil, error
	}

	var dto DTO
	mapper.Map(&dto, updatedModel)

	return &dto, nil
}

func (service *DeleteServiceImpl[Model, ID, IDRequest]) Delete(idRequest IDRequest) error {
	var id ID
	mapper.Map(&id, &idRequest)
	return service.repository.Delete(id)
}

func (service *ReadServiceImpl[Model, ID, IDRequest, DTO]) FindOne(idRequest IDRequest) (*DTO, error) {
	var id ID
	mapper.Map(&id, &idRequest)

	model, error := service.repository.FindById(id)
	if error != nil {
		return nil, error
	}
	var dto DTO
	mapper.Map(&dto, model)
	return &dto, nil
}

func (service *ReadServiceImpl[Model, ID, IDRequest, DTO]) FindAll() ([]*DTO, error) {
	models, error := service.repository.FindAll()
	if error != nil {
		return nil, error
	}
	var dtos []*DTO
	mapper.Map(&dtos, models)
	return dtos, nil
}
