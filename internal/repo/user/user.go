package user

import (
	"xm/internal/entity/models"

	"gorm.io/gorm"
)

type RepoInterface interface {
	CreateUser(user *models.User) (result *models.User, err error)
	GetUser(name, password string) (result *models.User, err error)
}

type repo struct {
	db *gorm.DB
}

func InitUserRepo(db *gorm.DB) RepoInterface {
	return &repo{
		db: db,
	}
}
