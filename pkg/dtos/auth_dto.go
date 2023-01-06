package dtos

import "time"

type accountInfo struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type userInfo struct {
	Firstname string `json:"firstName"`
	Lastname string `json:"lastName"`
	Gender bool `json:"gender"`
	Birthdate time.Time `json:"birthDate"`
}

type RegisterRequest struct {
	Account accountInfo `json:"account"`
	User userInfo `json:"user"`
}

type LoginRequest struct {
	accountInfo
}


type UserDTO struct {
	ID uint
	userInfo
	Role string
}
