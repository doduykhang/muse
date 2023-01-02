package utils

import (
	"github.com/doduykhang/muse/pkg/service"
	"github.com/gorilla/schema"
)

var (
	decoder = schema.NewDecoder()
	componentService = service.GetComponentServices()
)
