package lib

import (
	"encoding/json"
	"net/http"

	"github.com/doduykhang/muse/pkg/service"
	"github.com/doduykhang/muse/pkg/utils"
	"github.com/go-chi/chi/v5"
)

var (
	componentService = service.GetComponentServices()
)

// Experimental
type (
	MyChi                   chi.Mux
	RequestType             = func(w http.ResponseWriter, r *http.Request)
	MyRequest[T any, R any] func(body *T) (*R, error)
)

type HttpRequest[T, R any] struct {
	method   string
	endpoint string
}

func (request *HttpRequest[T, R]) Get(endpoint string) *HttpRequest[T, R] {
	request.method = "GET"
	request.endpoint = endpoint

	return request
}

func Get[T, R any](endpoint string) func(r chi.Router, handleFunc MyRequest[T, R]) {
	return func(r chi.Router, handleFunc MyRequest[T, R]) {
		r.Get(endpoint, Request(handleFunc))
	}
}

func (request *HttpRequest[T, R]) HandleFunc(r chi.Router, handleFunc MyRequest[T, R]) {
	r.With(
		utils.GetRequestBody[T],
	).MethodFunc(request.method, request.endpoint, Request(handleFunc))
}

func RequestTest[DTO any](r chi.Mux) chi.Router {
	return r.With(
		utils.GetRequestBody[DTO],
	)
}

func Request[T any, R any](request MyRequest[T, R]) RequestType {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, ok := ctx.Value("request").(*T)
		if !ok {
			utils.ErrorResponse(&w, "error getting body", 500)
			return
		}
		response, err := request(body)
		if err != nil {
			utils.ErrorResponse(&w, err.Error(), 500)
			return
		}
		utils.JsonResponse(&w, response)
	}

}

func LoginRequest[T any, R any](request MyRequest[T, R]) RequestType {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		body, ok := ctx.Value("request").(*T)
		if !ok {
			utils.ErrorResponse(&w, "error getting body", 500)
			return
		}
		response, err := request(body)
		if err != nil {
			utils.ErrorResponse(&w, err.Error(), 500)
			return
		}

		cookie := componentService.CookieService.CreateCookie()

		cacheValue, err := json.Marshal(response)
		if err != nil {
			utils.ErrorResponse(&w, err.Error(), 500)
			return
		}

		err = componentService.CacheService.Set(cookie.Value, string(cacheValue))

		if err != nil {
			utils.ErrorResponse(&w, err.Error(), 500)
			return
		}
		
		http.SetCookie(w, &cookie)
		utils.JsonResponse(&w, response)
	}
}

func JsonRequest[T any](r chi.Router, method string, pattern string, handleFunc RequestType, middlewares ...utils.MiddleWare) {
	middlewares = append(middlewares, utils.GetRequestBody[T])
	r.With(
		middlewares...,
	).MethodFunc(method, pattern, handleFunc)
}
