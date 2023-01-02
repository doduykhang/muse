package service

import "golang.org/x/crypto/bcrypt"

type PasswordService interface {
	Generate(rawPassword string) (string, error) 
	Compare(password string, myPassword string) bool
}

type passwordService struct {

}

func (service *passwordService) Generate(rawPassword string) (string, error) {
    	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
    	if err != nil {
		return "", err
    	}
	return string(hashedPassword), nil
} 

func (service *passwordService) Compare(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
} 

func NewPasswordService() PasswordService {
	return &passwordService{}
}
