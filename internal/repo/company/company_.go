package company

import (
	"time"
	"xm/internal/entity/models"
	"xm/logger"

	"gorm.io/gorm/clause"
)

func (d *repo) CreateCompanyData(company *models.Company) (result *models.Company, err error) {
	logger.Info("Begin Repo - CreateCompanyData")

	err = d.db.Table("company").Create(&company).Error
	if err != nil {
		logger.Error("error creating company", err.Error())
	}

	logger.Info("End Repo - CreateCompanyData", company.ID)
	return company, err
}

func (d *repo) GetCompanyData(id string) (result *models.Company, err error) {
	logger.Info("Begin Repo - GetCompanyData")

	err = d.db.Table("company").Where("id = ?", id).First(&result).Error
	if err != nil {
		logger.Error("error fetching from company", err.Error())
	}

	logger.Info("End Repo - GetCompanyData", result)
	return result, err
}

func (d *repo) PatchCompanyData(id string, obj *models.Company) (result *models.Company, err error) {
	logger.Info("Begin Repo - PatchCompanyData")

	updateMap := make(map[string]interface{}, 5)
	if len(obj.Name) > 0 {
		updateMap["name"] = obj.Name
	}
	if len(obj.Description) > 0 {
		updateMap["description"] = obj.Description
	}
	if obj.EmployeeCount != nil {
		updateMap["employee_count"] = obj.EmployeeCount
	}
	if obj.IsRegistered != nil {
		updateMap["is_registered"] = obj.IsRegistered
	}
	if len(obj.Type) > 0 {
		updateMap["type"] = obj.Type
	}

	if len(updateMap) > 0 {
		updateMap["updated_at"] = time.Now().UTC()

		result = &models.Company{ID: id}
		if err = d.db.Table("company").Model(&result).Clauses(clause.Returning{}).Updates(updateMap).Error; err != nil {
			logger.Error("error fetching from company", err.Error())
		}
	}

	logger.Info("End Repo - PatchCompanyData", result)
	return result, err
}

func (d *repo) DeleteCompanyData(id string) (err error) {
	logger.Info("Begin Repo - DeleteCompanyData")

	err = d.db.Table("company").Delete(&models.Company{}, "id = ?", id).Error
	if err != nil {
		logger.Error("error fetching from company", err.Error())
	}

	logger.Info("End Repo - DeleteCompanyData")
	return err
}
