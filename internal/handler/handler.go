package handler

import (
	"xm/internal/usecase"
)

type Handler struct {
	CompanyHandler *companyHandler
	UserHandler    *userHandler
	HealthHandler  *healthHandler
}

func Init(uc *usecase.Usecase) *Handler {
	return &Handler{
		CompanyHandler: InitCompanyHandler(uc.Company),
		UserHandler:    InitUserHandler(uc.User),
		HealthHandler:  InitHealthHandler(),
	}
}
