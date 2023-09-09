package user

import (
	"xm/internal/entity/models"
	"xm/logger"
)

func (d *repo) GetUser(name, password string) (result *models.User, err error) {
	logger.Info("Begin Repo - GetUser")

	err = d.db.Table("users").Where("name = ? AND password = ?", name, password).First(&result).Error
	if err != nil {
		logger.Error("error fetching from company", err.Error())
	}

	logger.Info("End Repo - GetUser", result)
	return result, err
}

func (d *repo) CreateUser(user *models.User) (result *models.User, err error) {
	logger.Info("Begin Repo - CreateUser")

	err = d.db.Table("users").Create(&user).Error
	if err != nil {
		logger.Error("error creating company", err.Error())
	}

	logger.Info("End Repo - CreateUser", user.ID)
	return user, err
}
