package handler

import (
	"github.com/CristhianRamirez3010/task-manager-go/src/config/responseDto"
	"github.com/CristhianRamirez3010/task-manager-go/src/v1/handler/impl"
)

type IUserHandler interface {
	GetDocuments() responseDto.DesponseDto
}

func BuildIUserHandler() IUserHandler {
	return impl.BuildUserHandler()
}
