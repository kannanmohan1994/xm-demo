package usecase

import (
	"xm/internal/middleware"
	"xm/internal/repo"
	company "xm/internal/usecase/company"
	"xm/internal/usecase/user"
)

type Usecase struct {
	Company company.UsecaseInterface
	User    user.UsecaseInterface
}

func Init(repo *repo.Repo, middleware middleware.Middleware) *Usecase {
	return &Usecase{
		Company: company.InitCompanyUsecase(repo.Company),
		User:    user.InitUserUsecase(repo.User, middleware),
	}
}
