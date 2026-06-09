package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mizanalyst/mizanalyst/pagination"
)

// OK sends a 200 response with no data.
func OK(c *gin.Context, message string) {
	resp := NewResponse(true, message, nil, http.StatusOK)
	c.JSON(http.StatusOK, resp.GetBody())
}

// OKWithBody sends a 200 response with a data payload.
func OKWithBody(c *gin.Context, message string, data interface{}) {
	resp := NewResponse(true, message, data, http.StatusOK)
	c.JSON(http.StatusOK, resp.GetBody())
}

// BadRequest sends a 400 response with no data.
func BadRequest(c *gin.Context, message string) {
	resp := NewResponse(false, message, nil, http.StatusBadRequest)
	c.JSON(http.StatusBadRequest, resp.GetBody())
}

// BadRequestWithBody sends a 400 response with a data payload (e.g., validation errors).
func BadRequestWithBody(c *gin.Context, message string, data interface{}) {
	resp := NewResponse(false, message, data, http.StatusBadRequest)
	c.JSON(http.StatusBadRequest, resp.GetBody())
}

// Unauthorized sends a 401 response with no data.
func Unauthorized(c *gin.Context, message string) {
	resp := NewResponse(false, message, nil, http.StatusUnauthorized)
	c.JSON(http.StatusUnauthorized, resp.GetBody())
}

// UnauthorizedWithBody sends a 401 response with a data payload.
func UnauthorizedWithBody(c *gin.Context, message string, data interface{}) {
	resp := NewResponse(false, message, data, http.StatusUnauthorized)
	c.JSON(http.StatusUnauthorized, resp.GetBody())
}

// OKPaginated sends a 200 response with a paginated payload.
func OKPaginated(c *gin.Context, message string, data interface{}, meta pagination.PaginationMeta) {
	payload := pagination.PaginatedResult{
		Data: data,
		Meta: meta,
	}
	resp := NewResponse(true, message, payload, http.StatusOK)
	c.JSON(http.StatusOK, resp.GetBody())
}
