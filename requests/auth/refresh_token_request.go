package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mizanalyst/mizanalyst/dtos"
	"github.com/mizanalyst/mizanalyst/requests"
)

// RefreshTokenRequest handles binding and validation for the refresh token endpoint.
type RefreshTokenRequest struct {
	Body dtos.RefreshTokenDTO
}

// Run executes the request chain: Bind → Validate.
// Returns false if any step fails (response is already written to context).
func (r *RefreshTokenRequest) Run(c *gin.Context) bool {
	return requests.RunChain(c,
		&requests.RequestBinder{Target: &r.Body},
		&requests.RequestValidator{Target: &r.Body},
	)
}
