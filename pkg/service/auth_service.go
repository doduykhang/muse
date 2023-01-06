package service

import (
	"errors"
	"fmt"

	"github.com/doduykhang/muse/pkg/constants"
	"github.com/doduykhang/muse/pkg/dtos"
	"github.com/doduykhang/muse/pkg/models"
	"github.com/doduykhang/muse/pkg/repositories"
)

type (
	RegisterRequest = dtos.RegisterRequest
	LoginRequest = dtos.LoginRequest
	UserDTO = dtos.UserDTO
)

type AuthService interface {
	Register(request *RegisterRequest) (*UserDTO, error)
	Login(request *LoginRequest) (*UserDTO, error)
}

func NewAuthService() AuthService {
	return &authService{
		accountRepository: appRepositores.AccountRepository,
		userRepository: appRepositores.UserRepository,
		passwordService: componentService.PasswordService,
	}
}

type authService struct{
	accountRepository repositories.AccountRepository
	userRepository repositories.UserRepository
	passwordService PasswordService
}

func (service *authService) Register(request *RegisterRequest) (*UserDTO, error) {
	hashPassword, err := service.passwordService.Generate(request.Account.Password) 

	if err != nil {
		return nil, err
	}

	var account *models.Account = &models.Account{}
	mapper.Map(account, &request.Account)
	account.Password = hashPassword
	account.Role = constants.USER_ROLE
	
	account, err = service.accountRepository.Create(account)
	if err != nil {
		return nil, err
	}

	var user *models.User = &models.User{}
	mapper.Map(user, &request.User)
	user.AccountID = account.ID 
	user, err = service.userRepository.Create(user)	

	if err != nil {
		return nil, err
	}

	var userDTO dtos.UserDTO
	mapper.Map(&userDTO, user)
	
	return &userDTO, nil
}

func (service *authService) Login(request *LoginRequest) (*UserDTO, error) {
	account, err := service.accountRepository.FindAccountByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	fmt.Println("request", request)

	ok := service.passwordService.Compare(request.Password, account.Password)
	if !ok {
		return nil, errors.New("wrong email or password")
	}

	user, err := service.userRepository.FindUserByAccountID(account.ID)
	if err != nil {
		return nil, err
	}
	
	var userDTO dtos.UserDTO
	mapper.Map(&userDTO, user)

	userDTO.Role = account.Role

	return &userDTO, nil
}
