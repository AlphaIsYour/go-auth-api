package main

import (
	"go-auth-api/controllers"
	"go-auth-api/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/profile", controllers.Profile)

	r.Run(":8080")
}
