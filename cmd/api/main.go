package main

import (
	"go-project/routes"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    // Create a new Gin router
    r := gin.Default()

    // Hello World Route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "It's work!",
		})
	})

    // Setup routes
    routes.SetupUserRoutes(r)

    // Start the server
    log.Println("Server is running on port 8090...")
    if err := r.Run(":8090"); err != nil {
        log.Fatalf("Failed to start the server: %v", err)
        return
    }
}
