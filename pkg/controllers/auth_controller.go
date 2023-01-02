package controllers

import (
	"github.com/doduykhang/muse/pkg/lib"
)


var (
	RegisterController = lib.Request(businessServices.AuthService.Register)
	LoginController = lib.LoginRequest(businessServices.AuthService.Login)
)
 
