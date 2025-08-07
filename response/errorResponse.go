package response

import "github.com/gin-gonic/gin"

// ErrorResponse represents a standardized error response format
type ErrorResponse struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
}

// SendError sends a standardized error response
func Error(c *gin.Context, message string, statusCode int) {
    c.JSON(statusCode, ErrorResponse{
        Code:    statusCode,
        Message: message,
    })
}
