package main

import (
	"go-auth-api/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected route (nanti pakai middleware JWT)
	r.GET("/profile", controllers.Profile)

	r.Run(":8080") // jalanin di http://localhost:8080
}
