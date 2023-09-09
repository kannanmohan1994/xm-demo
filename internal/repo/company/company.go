package company

import (
	"xm/internal/entity/models"

	"gorm.io/gorm"
)

type RepoInterface interface {
	GetCompanyData(id string) (result *models.Company, err error)
	CreateCompanyData(obj *models.Company) (result *models.Company, err error)
	PatchCompanyData(id string, obj *models.Company) (result *models.Company, err error)
	DeleteCompanyData(id string) (err error)
}

type repo struct {
	db *gorm.DB
}

func InitCompanyRepo(db *gorm.DB) RepoInterface {
	return &repo{
		db: db,
	}
}
