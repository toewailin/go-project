package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SuccessResponse represents a standardized success response format
type SuccessResponse struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}

// SendSuccess sends a standardized success response
func Success(c *gin.Context, data interface{}, message string) {
    c.JSON(http.StatusOK, SuccessResponse{
        Code:    http.StatusOK,
        Message: message,
        Data:    data,
    })
}
