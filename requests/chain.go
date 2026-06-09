package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/mizanalyst/mizanalyst/responses"
)

// ChainStep is the interface for each step in the request handling chain.
// Handle returns:
//   - ok: true if the step succeeded and the chain should continue
//   - error message is written to gin.Context if the step fails
type ChainStep interface {
	Handle(c *gin.Context) bool
}

// RunChain executes the given chain steps in order. If any step returns false
// (indicating it has already written an error response), the chain halts immediately.
func RunChain(c *gin.Context, steps ...ChainStep) bool {
	for _, step := range steps {
		if !step.Handle(c) {
			return false
		}
	}
	return true
}

// RequestBinder binds the incoming JSON payload to the target struct.
type RequestBinder struct {
	Target interface{}
}

// Handle performs JSON binding. Returns false and writes a 400 response on failure.
func (rb *RequestBinder) Handle(c *gin.Context) bool {
	if err := c.ShouldBindJSON(rb.Target); err != nil {
		responses.BadRequestWithBody(c, "Invalid request payload", gin.H{"error": err.Error()})
		return false
	}
	return true
}

// RequestValidator validates the bound struct using the application's validator.
type RequestValidator struct {
	Target interface{}
}

// Handle performs struct validation. Returns false and writes a 400 response on failure.
func (rv *RequestValidator) Handle(c *gin.Context) bool {
	if errs := Validate(rv.Target); len(errs) > 0 {
		responses.BadRequestWithBody(c, "Validation failed", gin.H{"errors": errs})
		return false
	}
	return true
}

// RequestUserIdRetriever extracts the user ID from the gin context (set by auth middleware).
type RequestUserIdRetriever struct {
	UserID *uint
}

// Handle retrieves the user ID from the context. Returns false and writes a 401 response if absent.
func (ru *RequestUserIdRetriever) Handle(c *gin.Context) bool {
	id, exists := c.Get("user_id")
	if !exists {
		responses.Unauthorized(c, "User not authenticated")
		return false
	}

	userID, ok := id.(uint)
	if !ok {
		responses.Unauthorized(c, "Invalid user identity")
		return false
	}

	*ru.UserID = userID
	return true
}
