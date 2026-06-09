package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/mizanalyst/mizanalyst/dtos"
	"github.com/mizanalyst/mizanalyst/requests"
)

// LoginRequest handles binding and validation for the login endpoint.
type LoginRequest struct {
	Body dtos.LoginDTO
}

// Run executes the request chain: Bind → Validate.
// Returns false if any step fails (response is already written to context).
func (r *LoginRequest) Run(c *gin.Context) bool {
	return requests.RunChain(c,
		&requests.RequestBinder{Target: &r.Body},
		&requests.RequestValidator{Target: &r.Body},
	)
}
