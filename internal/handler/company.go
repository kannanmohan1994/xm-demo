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

func (h *companyHandler) HandleCreateCompany(c *gin.Context) {
	logger.Info("Begin Handler - CreateCompany")

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

	res, err := h.companyUC.CreateCompany(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - CreateCompany")
	c.JSON(http.StatusOK, utils.Send(res))
}

func (h *companyHandler) HandleGetCompany(c *gin.Context) {
	logger.Info("Begin Handler - GetCompany")

	companyId := c.Param("company_id")

	result, err := h.companyUC.GetCompany(companyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - GetCompany")
	c.JSON(http.StatusOK, utils.Send(result))
}

func (h *companyHandler) HandlePatchCompany(c *gin.Context) {
	logger.Info("Begin Handler - PatchCompany")

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

	result, err := h.companyUC.PatchCompany(companyId, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - PatchCompany")
	c.JSON(http.StatusOK, utils.Send(result))
}

func (h *companyHandler) HandleDeleteCompany(c *gin.Context) {
	logger.Info("Begin Handler - DeleteCompany")

	companyId := c.Param("company_id")

	err := h.companyUC.DeleteCompany(companyId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - DeleteCompany")
	c.JSON(http.StatusOK, utils.Send(nil))
}
