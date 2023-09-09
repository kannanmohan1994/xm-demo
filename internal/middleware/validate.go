package middleware

import (
	"net/http"
	"xm/internal/utils"
	"xm/logger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func Validate[T any](tag string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestModel T
		err := c.BindJSON(&requestModel)
		if err != nil {
			logger.Errorf("error in BindJSON", err)
			c.IndentedJSON(http.StatusBadRequest, utils.Fail(100, err.Error()))
			c.Abort()
		}
		validate := validator.New()
		if err := validate.Struct(&requestModel); err != nil {
			logger.Errorf("error in struct validation", err)
			c.IndentedJSON(http.StatusBadRequest, utils.Fail(100, err.Error()))
			c.Abort()
		}
		c.Set(tag, requestModel)
	}
}
