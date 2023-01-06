package service

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

type CookieService interface {
	CreateCookie() http.Cookie	
}

type cookieService struct {}

func (service *cookieService) CreateCookie() http.Cookie {
	return http.Cookie{
		Name: 	"token",
		Path: "/",
		Value: 	uuid.New().String(),
		Expires: time.Now().Add(365 * 24 * time.Hour),
	}
}

func NewCookieService () CookieService {
	return &cookieService{}
}
