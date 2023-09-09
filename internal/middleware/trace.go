package middleware

import (
	"xm/consts"
	"xm/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// middleware that injects a 'X-Request-Id' into the context and header of each request
func (m *middleware) Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		xRequestID := uuid.New().String()
		c.Set(consts.CorrelationID, xRequestID)

		logger.Debugf(`[API-Hit] [X-Request-Id:%s] - "%s %s"`, xRequestID, c.Request.Method, c.Request.URL.Path)
		c.Next()
	}
}

// // GetCorrelationIDFromContext returns correlation-id from the given context if present
// func GetCorrelationIDFromContext(c *gin.Context) string {
// 	if v, ok := c.Get(consts.CorrelationID); ok {
// 		if requestID, ok := v.(string); ok {
// 			return requestID
// 		}
// 	}
// 	return ""
// }
