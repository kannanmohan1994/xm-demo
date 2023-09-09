package company

import (
	"time"
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	"xm/logger"

	"github.com/google/uuid"
)

func (u *usecase) CreateCompanyData(req *request.CreateCompanyRequest) (result *models.Company, err error) {
	logger.Info("Begin Usecase - CreateCompanyData")
	company := &models.Company{
		ID:            uuid.New().String(),
		Name:          req.Name,
		Description:   req.Description,
		EmployeeCount: req.EmployeeCount,
		IsRegistered:  req.IsRegistered,
		Type:          req.Type,
		CreatedAt:     time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
	}
	result, err = u.company.CreateCompanyData(company)
	if err != nil {
		return nil, err
	}
	logger.Info("End Usecase - CreateCompanyData")
	return result, err
}

func (u *usecase) GetCompanyData(id string) (result *models.Company, err error) {
	logger.Info("Begin Usecase - GetCompanyData")
	result, err = u.company.GetCompanyData(id)
	if err != nil {
		return result, err
	}
	logger.Info("End Usecase - GetCompanyData")
	return result, nil
}

func (u *usecase) PatchCompanyData(id string, req *request.PatchCompanyRequest) (result *models.Company, err error) {
	logger.Info("Begin Usecase - PatchCompanyData")

	result, err = u.company.GetCompanyData(id)
	if err != nil {
		return result, err
	}

	company := &models.Company{
		Name:          req.Name,
		Description:   req.Description,
		EmployeeCount: req.EmployeeCount,
		IsRegistered:  req.IsRegistered,
		Type:          req.Type,
	}
	result, err = u.company.PatchCompanyData(id, company)
	if err != nil {
		return result, err
	}
	logger.Info("End Usecase - PatchCompanyData")
	return result, nil
}

func (u *usecase) DeleteCompanyData(id string) (err error) {
	logger.Info("Begin Usecase - PatchCompanyData")

	_, err = u.company.GetCompanyData(id)
	if err != nil {
		return err
	}

	err = u.company.DeleteCompanyData(id)
	if err != nil {
		return err
	}

	logger.Info("End Usecase - PatchCompanyData")
	return nil
}
