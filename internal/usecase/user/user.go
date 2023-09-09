package user

import (
	"xm/internal/entity/response"
	"xm/internal/middleware"
	user "xm/internal/repo/user"
	"xm/logger"
)

type UsecaseInterface interface {
	CreateUser(name, password string) (resp *response.UserResponse, err error)
	LoginUser(name, password string) (resp *response.UserResponse, err error)
}

type usecase struct {
	logger     logger.Log
	middleware middleware.Middleware
	user       user.UserRepository
}

func InitUserUsecase(user user.UserRepository, middleware middleware.Middleware, logger logger.Log) UsecaseInterface {
	return &usecase{
		logger:     logger,
		middleware: middleware,
		user:       user,
	}
}
