package handler

import (
	"net/http"
	"xm/internal/entity/request"
	company "xm/internal/usecase/company"
	"xm/internal/utils"
	"xm/logger"

	"github.com/gin-gonic/gin"
)

type companyHandler struct {
	companyUC company.UsecaseInterface
}

func InitCompanyHandler(uc company.UsecaseInterface) *companyHandler {
	return &companyHandler{
		companyUC: uc,
	}
}

func (h *companyHandler) HandleCreateCompanyData(c *gin.Context) {
	logger.Info("Begin Handler - CreateCompanyData")

	if errMessage, ok := c.Get("error"); ok {
		c.JSON(http.StatusBadRequest, utils.Fail(100, errMessage.(string)))
		return
	}

	data, ok := c.Get("CreateCompanyRequest")
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(100, utils.ErrInvalidRequest.Error()))
		return
	}
	req := data.(request.CreateCompanyRequest)

	res, err := h.companyUC.CreateCompanyData(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - CreateCompanyData")
	c.JSON(http.StatusOK, utils.Send(res))
}

func (h *companyHandler) HandleGetCompanyData(c *gin.Context) {
	logger.Info("Begin Handler - GetCompanyData")

	companyId := c.Param("company_id")

	result, err := h.companyUC.GetCompanyData(companyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - GetCompanyData")
	c.JSON(http.StatusOK, utils.Send(result))
}

func (h *companyHandler) HandlePatchCompanyData(c *gin.Context) {
	logger.Info("Begin Handler - PatchCompanyData")

	if errMessage, ok := c.Get("error"); ok {
		c.JSON(http.StatusBadRequest, utils.Fail(100, errMessage.(string)))
		return
	}

	data, ok := c.Get("PatchCompanyRequest")
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(100, utils.ErrInvalidRequest.Error()))
		return
	}
	req := data.(request.PatchCompanyRequest)
	companyId := c.Param("company_id")

	result, err := h.companyUC.PatchCompanyData(companyId, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - PatchCompanyData")
	c.JSON(http.StatusOK, utils.Send(result))
}

func (h *companyHandler) HandleDeleteCompanyData(c *gin.Context) {
	logger.Info("Begin Handler - DeleteCompanyData")

	companyId := c.Param("company_id")

	err := h.companyUC.DeleteCompanyData(companyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - DeleteCompanyData")
	c.JSON(http.StatusOK, utils.Send(nil))
}
