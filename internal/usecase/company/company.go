package company

import (
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	company "xm/internal/repo/company"
)

type UsecaseInterface interface {
	GetCompanyData(id string) (result *models.Company, err error)
	CreateCompanyData(req *request.CreateCompanyRequest) (result *models.Company, err error)
	PatchCompanyData(id string, req *request.PatchCompanyRequest) (result *models.Company, err error)
	DeleteCompanyData(id string) (err error)
}

type usecase struct {
	company company.RepoInterface
}

func InitCompanyUsecase(company company.RepoInterface) UsecaseInterface {
	return &usecase{
		company: company,
	}
}
