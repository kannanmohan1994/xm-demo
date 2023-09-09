package repo

import (
	company "xm/internal/repo/company"
	"xm/internal/repo/user"

	"gorm.io/gorm"
)

type Repo struct {
	Company company.RepoInterface
	User    user.RepoInterface
}

func InitRepo(db *gorm.DB) *Repo {
	return &Repo{
		Company: company.InitCompanyRepo(db),
		User:    user.InitUserRepo(db),
	}
}
