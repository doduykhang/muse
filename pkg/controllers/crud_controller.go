package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/service"
	"github.com/doduykhang/muse/pkg/utils"
	"github.com/dranikpg/dto-mapper"
	"github.com/gorilla/schema"
)

var (
	decoder = schema.NewDecoder()
	mapper  = dto.Mapper{}
)

type CreateController[Request any, DTO any] struct {
	service service.CreateService[Request, DTO]
}

type UpdateController[Request any, DTO any, ID any] struct {
	service service.UpdateService[Request, DTO, ID]
}

type DeleteController[ID any] struct {
	service service.DeleteService[ID]
}

type ReadController[DTO any, ID any] struct {
	service service.ReadService[DTO, ID]
}

type CreateUploadController[Request any, DTO any] struct {
	service service.CreateUploadService[Request, DTO]
}

func (controller CreateController[Request, DTO]) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	ctx := r.Context()
	request, ok := ctx.Value("request").(*Request)
	if !ok {
		utils.ErrorResponse(&w, "error getting request", 500)
		return
	}

	dto, err := controller.service.Create(request)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(response)
}

func (controller UpdateController[Request, DTO, ID]) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	ctx := r.Context()
	request, ok := ctx.Value("request").(*Request)
	if !ok {
		utils.ErrorResponse(&w, "error getting request", 500)
		return
	}

	id, ok := ctx.Value("id").(*ID)
	if !ok {
		utils.ErrorResponse(&w, "error getting id", 500)
		return
	}

	dto, err := controller.service.Update(*id, request)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(response)
}

func (controller DeleteController[ID]) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	ctx := r.Context()
	id, ok := ctx.Value("id").(*ID)
	if !ok {
		utils.ErrorResponse(&w, "error getting id", 500)
		return
	}
	err := controller.service.Delete(*id)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response, err := json.Marshal("ok")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(response)
}

func (controller ReadController[DTO, ID]) FindOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	ctx := r.Context()
	id, ok := ctx.Value("id").(*ID)
	if !ok {
		utils.ErrorResponse(&w, "error getting id", 500)
		return
	}
	dto, err := controller.service.FindOne(*id)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(response)
}

func (controller ReadController[DTO, ID]) FindAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	dto, err := controller.service.FindAll()

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	response, err := json.Marshal(dto)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Write(response)
}

func (controller CreateUploadController[Request, DTO]) CreateUpload(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	files, ok := ctx.Value("fileList").([]*dtos.FileDTO)
	if !ok {
		utils.ErrorResponse(&w, "error getting files", 500)
		return
	}

	request, ok := ctx.Value("postForm").(*Request)
	if !ok {
		utils.ErrorResponse(&w, "error getting postForm", 500)
		return
	}

	dto, err := controller.service.CreateUpload(request, files)
	if err != nil {
		utils.ErrorResponse(&w, err.Error(), 500)
		return
	}

	utils.JsonResponse(&w, dto)
}
