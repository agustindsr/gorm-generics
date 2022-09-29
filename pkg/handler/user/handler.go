package user

import (
	"gorm-with-generics/pkg/repository/users"
	"gorm-with-generics/pkg/service/user"
)

type handler struct {
	Service user.Service
}

func New() handler {
	userRepo := users.NewRepository()

	return handler{
		Service: user.NewService(userRepo),
	}
}
