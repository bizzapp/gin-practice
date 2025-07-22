package response

import (
	"github.com/gin-gonic/gin"
)

type ErrorDetail struct {
	Code    string `json:"code"`
	Details string `json:"details"`
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *ErrorDetail `json:"error,omitempty"`
}

// Success response
func JSON(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error response
func Error(c *gin.Context, status int, code string, message string, detail string) {
	c.JSON(status, Response{
		Success: false,
		Message: message,
		Error: &ErrorDetail{
			Code:    code,
			Details: detail,
		},
	})
}
