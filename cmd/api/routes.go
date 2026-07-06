package main

import (
	"github.com/deeep8250/vibecheck-api/internal/auth"
	"github.com/deeep8250/vibecheck-api/internal/feed"
	"github.com/deeep8250/vibecheck-api/internal/follow"
	"github.com/deeep8250/vibecheck-api/internal/middleware"
	"github.com/deeep8250/vibecheck-api/internal/post"
	"github.com/deeep8250/vibecheck-api/internal/react"
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
	v1.GET("/posts", middleware.Middleware(), PostHandler.GetAllPostHandler)
	v1.PATCH("/post/:id", middleware.Middleware(), PostHandler.UpdatePostHandler)
	v1.DELETE("/post/:id", middleware.Middleware(), PostHandler.DeletePostHandler)

	//follow
	followRepo := follow.NewFollowRepository()
	followService := follow.NewFollowService(followRepo)
	followHandler := follow.NewFollowHand(followService)

	v1.POST("/follow/:id", middleware.Middleware(), followHandler.FollowHandler)
	v1.DELETE("/unfollow/:id", middleware.Middleware(), followHandler.UnfollowHandler)
	v1.GET("/follow", middleware.Middleware(), followHandler.GetFollowHandler)

	//feed
	feedRepo := feed.NewFeedRepo()
	feedService := feed.NewFeedService(feedRepo)
	feedHandler := feed.NewFeedHandler(feedService)

	v1.GET("/feed", middleware.Middleware(), feedHandler.GetFeed)

	//react
	reactRepo := react.NewReactRepository()
	reactService := react.NewReactService(reactRepo)
	reactHandler := react.NewReactHandler(reactService)

	v1.POST("/react/:id", middleware.Middleware(), reactHandler.ReactPost)
	v1.GET("/reactions/:id", reactHandler.GetReactionByPost)
	v1.DELETE("/reaction/:postid", middleware.Middleware(), reactHandler.DeleteReaction)

	r.Run(":8080")

}
