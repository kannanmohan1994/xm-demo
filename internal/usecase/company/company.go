package company

import (
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	company "xm/internal/repo/company"
	"xm/logger"
)

type UsecaseInterface interface {
	GetCompany(id string) (result *models.Company, err error)
	CreateCompany(req *request.CreateCompanyRequest) (result *models.Company, err error)
	PatchCompany(id string, req *request.PatchCompanyRequest) (result *models.Company, err error)
	DeleteCompany(id string) (err error)
}

type usecase struct {
	logger  logger.Log
	company company.CompanyRepository
}

func InitCompanyUsecase(company company.CompanyRepository, logger logger.Log) UsecaseInterface {
	return &usecase{
		logger:  logger,
		company: company,
	}
}
