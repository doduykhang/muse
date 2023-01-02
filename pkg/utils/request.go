package utils

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/doduykhang/muse/pkg/dtos"
)

type (
	MiddleWare = func(http.Handler) http.Handler
)

func GetRequestBody[Request any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var request Request
		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			ErrorResponse(&w, err.Error(), 400)
			return
		}

		ctx := context.WithValue(r.Context(), "request", &request)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestForm[Request any](next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(10 << 20)
		if err != nil {
			ErrorResponse(&w, err.Error(), 400)
			return
		}

		var request Request
		err = decoder.Decode(&request, r.PostForm)
		if err != nil {
			ErrorResponse(&w, err.Error(), 400)
			return
		}
		ctx := context.WithValue(r.Context(), "postForm", &request)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetRequestQuery[Request any](name string) MiddleWare {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var request Request
			err := decoder.Decode(&request, r.URL.Query())
			if err != nil {
				ErrorResponse(&w, err.Error(), 400)
				return
			}
			ctx := context.WithValue(r.Context(), name, &request)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetFormFile(fileName string, field string, path string) MiddleWare {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r.ParseMultipartForm(10 << 20)

			file, fileHandler, err := r.FormFile(fileName)
			if err != nil {
				ErrorResponse(&w, err.Error(), 500)
			}

			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				ErrorResponse(&w, err.Error(), 500)
			}

			fileDto := &dtos.FileDTO{
				Bytes: fileBytes,
				Name:  fileHandler.Filename,
				Size:  fileHandler.Size,
				Field: field,
				Path:  path,
			}

			ctx := r.Context()
			fileList, ok := ctx.Value("fileList").([]*dtos.FileDTO)
			if !ok {
				ctx = context.WithValue(r.Context(), "fileList", []*dtos.FileDTO{fileDto})
			} else {
				ctx = context.WithValue(r.Context(), "fileList", append(fileList, fileDto))
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
