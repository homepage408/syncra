package response

// shared/response: standard API response helpers

type SuccessResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Errors  interface{} `json:"errors,omitempty"`
}
