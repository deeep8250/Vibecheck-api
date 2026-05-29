package main

import (
	"github.com/deeep8250/vibecheck-api/internal/auth"
	"github.com/deeep8250/vibecheck-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	v1 := r.Group("/api/v1", middleware.ErrorHandle())

	//dependency injection

	//auth
	authRepo := auth.NewDb()
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	v1.POST("/register", authHandler.RegisterUser)
	v1.POST("/login", authHandler.Login)

	r.Run(":8080")

}
