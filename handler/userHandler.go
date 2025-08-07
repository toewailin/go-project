package handler

import (
	"go-project/models"
	"go-project/response"
	"go-project/serialization"

	"github.com/gin-gonic/gin"
)

// GetUser fetches a user and returns it as a response
func GetUser(c *gin.Context) {
    // For demonstration, creating a mock user
    user := models.User{
        ID:    "1",
        Name:  "John Doe",
        Email: "john@example.com",
    }

    // Serialize the user to DTO format
    userDTO := serialization.UserToDTO(user)

    // Send success response
    response.Success(c, userDTO, "User retrieved successfully")
}
