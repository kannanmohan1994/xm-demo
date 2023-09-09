package company

import (
	"xm/internal/entity/models"

	"gorm.io/gorm"
)

type RepoInterface interface {
	GetCompany(id string) (result *models.Company, err error)
	CreateCompany(obj *models.Company) (result *models.Company, err error)
	PatchCompany(id string, obj *models.Company) (result *models.Company, err error)
	DeleteCompany(id string) (err error)
}

type repo struct {
	db *gorm.DB
}

func InitCompanyRepo(db *gorm.DB) RepoInterface {
	return &repo{
		db: db,
	}
}
