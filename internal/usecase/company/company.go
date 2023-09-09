package company

import (
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	company "xm/internal/repo/company"
)

type UsecaseInterface interface {
	GetCompany(id string) (result *models.Company, err error)
	CreateCompany(req *request.CreateCompanyRequest) (result *models.Company, err error)
	PatchCompany(id string, req *request.PatchCompanyRequest) (result *models.Company, err error)
	DeleteCompany(id string) (err error)
}

type usecase struct {
	company company.RepoInterface
}

func InitCompanyUsecase(company company.RepoInterface) UsecaseInterface {
	return &usecase{
		company: company,
	}
}
