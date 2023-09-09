package docs

import (
	"xm/internal/entity/models"
	"xm/internal/entity/request"
)

// swagger:route POST /v1/company Company idCreateCompany
// Creates a Company
// responses:
// 		200: CreateCompanyResponseWrapper
//		400: CustomErrorWrapper

// swagger:parameters idCreateCompany
type CreateCompanyRequestWrapper struct {
	// in:body
	Body request.CreateCompanyRequest
}

// swagger:response CreateCompanyResponseWrapper
type CreateCompanyResponseWrapper struct {
	// in:body
	Body models.Company
}

// swagger:route GET /v1/company Company idGetCompany
// Fetch company data
// responses:
// 		200: GetCompanyResponseWrapper
//		400: CustomErrorWrapper

// swagger:response GetCompanyResponseWrapper
type GetCompanyResponseWrapper struct {
	// in:body
	Body models.Company
}
