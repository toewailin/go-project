package routes

import (
	"go-project/handler"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes defines routes for the user entity
func SetupUserRoutes(r *gin.Engine) {
    userRoutes := r.Group("/users")
    userRoutes.GET("/:id", handler.GetUser)
}
