package utils

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/doduykhang/muse/pkg/dtos"
)

func Authenticated() MiddleWare {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenCookie, err := r.Cookie("token")	
			if err != nil {
				ErrorResponse(&w, "who are you :<", 401)
				return
			}

			value, err := componentService.CacheService.Get(tokenCookie.Value)
			if err != nil {
				ErrorResponse(&w, "who are you :<", 401)
				return
			}

			var userDTO dtos.UserDTO
			err = json.Unmarshal([]byte(value), &userDTO)
			if err != nil {
				ErrorResponse(&w, err.Error(), 500)
				return
			}

			ctx := context.WithValue(r.Context(), "userId", userDTO.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Authorized(roles ...string) MiddleWare {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenCookie, err := r.Cookie("token")	
			if err != nil {
				ErrorResponse(&w, "who are you :<", 401)
				return
			}

			value, err := componentService.CacheService.Get(tokenCookie.Value)
			if err != nil {
				ErrorResponse(&w, "who are you :<", 401)
				return
			}

			var userDTO dtos.UserDTO
			err = json.Unmarshal([]byte(value), &userDTO)
			if err != nil {
				ErrorResponse(&w, err.Error(), 500)
				return
			}
			if !Contains(roles, userDTO.Role) {
				ErrorResponse(&w, "get the f out of here ::<<", 403)
				return
			}


			ctx := context.WithValue(r.Context(), "userId", userDTO.ID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
