package company

import (
	"time"
	"xm/internal/entity/models"
	"xm/internal/entity/request"
	"xm/logger"

	"github.com/google/uuid"
)

func (u *usecase) CreateCompany(req *request.CreateCompanyRequest) (result *models.Company, err error) {
	logger.Info("Begin Usecase - CreateCompany")
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
	result, err = u.company.CreateCompany(company)
	if err != nil {
		return nil, err
	}
	logger.Info("End Usecase - CreateCompany")
	return result, err
}

func (u *usecase) GetCompany(id string) (result *models.Company, err error) {
	logger.Info("Begin Usecase - GetCompany")
	result, err = u.company.GetCompany(id)
	if err != nil {
		return result, err
	}
	logger.Info("End Usecase - GetCompany")
	return result, nil
}

func (u *usecase) PatchCompany(id string, req *request.PatchCompanyRequest) (result *models.Company, err error) {
	logger.Info("Begin Usecase - PatchCompany")

	result, err = u.company.GetCompany(id)
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
	result, err = u.company.PatchCompany(id, company)
	if err != nil {
		return result, err
	}
	logger.Info("End Usecase - PatchCompany")
	return result, nil
}

func (u *usecase) DeleteCompany(id string) (err error) {
	logger.Info("Begin Usecase - PatchCompany")

	_, err = u.company.GetCompany(id)
	if err != nil {
		return err
	}

	err = u.company.DeleteCompany(id)
	if err != nil {
		return err
	}

	logger.Info("End Usecase - PatchCompany")
	return nil
}
