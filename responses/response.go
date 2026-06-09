package responses

// StandardBody is the unified JSON structure for all API responses.
type StandardBody struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Status  int         `json:"status"`
}

// IResponse is the interface that all response types must implement.
type IResponse interface {
	GetBody() StandardBody
}

// standardResponse implements IResponse.
type standardResponse struct {
	body StandardBody
}

// GetBody returns the StandardBody payload.
func (r *standardResponse) GetBody() StandardBody {
	return r.body
}

// NewResponse creates a new IResponse with the given parameters.
func NewResponse(success bool, message string, data interface{}, status int) IResponse {
	return &standardResponse{
		body: StandardBody{
			Success: success,
			Message: message,
			Data:    data,
			Status:  status,
		},
	}
}
