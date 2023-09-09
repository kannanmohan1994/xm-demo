package docs

import (
	"xm/internal/entity/response"
)

// swagger:route GET /health Health idGetHealth
// Checks the health of xm app
// responses:
// 		200: GetHealthResponse

// swagger:response GetHealthResponse
type GetHealthResponseWrapper struct {
	// in:body
	Body response.Health
}
