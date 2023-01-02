package utils

import (
	"encoding/json"
	"net/http"

	"github.com/doduykhang/muse/pkg/dtos"
)

func JsonResponse(w *http.ResponseWriter, x interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(x)
	(*w).WriteHeader(http.StatusOK)
	(*w).Write(res)
}

func ErrorResponse(w *http.ResponseWriter, message string, code int) {
	(*w).Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(
		dtos.ErrorResponse{
			Code:    code,
			Message: message,
		},
	)
	(*w).WriteHeader(code)
	(*w).Write(res)
}
