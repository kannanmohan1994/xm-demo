package handler

import (
	"net/http"
	"xm/internal/entity/request"
	user "xm/internal/usecase/user"
	"xm/internal/utils"
	"xm/logger"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userUC user.UsecaseInterface
}

func InitUserHandler(uc user.UsecaseInterface) *userHandler {
	return &userHandler{
		userUC: uc,
	}
}

func (h *userHandler) HandleCreateUser(c *gin.Context) {
	logger.Info("Begin Handler - CreateUser")

	if errMessage, ok := c.Get("error"); ok {
		c.JSON(http.StatusBadRequest, utils.Fail(100, errMessage.(string)))
		return
	}

	data, ok := c.Get("CreateUserRequest")
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(100, utils.ErrInvalidRequest.Error()))
		return
	}
	req := data.(request.UserRequest)

	res, err := h.userUC.CreateUser(req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - CreateCompany")
	c.JSON(http.StatusOK, utils.Send(res))
}

func (h *userHandler) HandleLoginUser(c *gin.Context) {
	logger.Info("Begin Handler - CreateUser")

	if errMessage, ok := c.Get("error"); ok {
		c.JSON(http.StatusBadRequest, utils.Fail(100, errMessage.(string)))
		return
	}

	data, ok := c.Get("LoginUserRequest")
	if !ok {
		c.JSON(http.StatusBadRequest, utils.Fail(100, utils.ErrInvalidRequest.Error()))
		return
	}
	req := data.(request.UserRequest)

	res, err := h.userUC.LoginUser(req.Name, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.Fail(100, err.Error()))
		return
	}

	logger.Info("End Handler - CreateCompany")
	c.JSON(http.StatusOK, utils.Send(res))
}
