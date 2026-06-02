package main

import (
	"github.com/deeep8250/vibecheck-api/internal/auth"
	"github.com/deeep8250/vibecheck-api/internal/middleware"
	"github.com/deeep8250/vibecheck-api/internal/post"
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

	//post
	postRepo := post.NewPostRepository()
	postService := post.NewPostService(postRepo)
	PostHandler := post.NewPostHandler(postService)

	v1.POST("/post", middleware.Middleware(), PostHandler.CreatePostHandler)
	v1.GET("/post/:id", middleware.Middleware(), PostHandler.GetPostHandler)
	v1.PATCH("/post/:id", middleware.Middleware(), PostHandler.UpdatePostHandler)
	v1.DELETE("/post/:id", middleware.Middleware(), PostHandler.DeletePostHandler)

	r.Run(":8080")

}
