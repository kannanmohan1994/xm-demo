package user

import (
	"xm/internal/entity/response"
	"xm/internal/middleware"
	user "xm/internal/repo/user"
)

type UsecaseInterface interface {
	CreateUser(name, password string) (resp *response.UserResponse, err error)
	LoginUser(name, password string) (resp *response.UserResponse, err error)
}

type usecase struct {
	middleware middleware.Middleware
	user       user.RepoInterface
}

func InitUserUsecase(user user.RepoInterface, middleware middleware.Middleware) UsecaseInterface {
	return &usecase{
		middleware: middleware,
		user:       user,
	}
}
