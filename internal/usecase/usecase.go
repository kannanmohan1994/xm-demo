package usecase

import (
	"xm/internal/middleware"
	"xm/internal/repo"
	company "xm/internal/usecase/company"
	"xm/internal/usecase/user"
	"xm/logger"
)

type Usecase struct {
	Company company.UsecaseInterface
	User    user.UsecaseInterface
}

func Init(repo *repo.Repo, middleware middleware.Middleware, logger logger.Log) *Usecase {
	return &Usecase{
		Company: company.InitCompanyUsecase(repo.Company, logger),
		User:    user.InitUserUsecase(repo.User, middleware, logger),
	}
}
