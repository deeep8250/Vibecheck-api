package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	r.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "alive",
		})
	})
	r.Run(":8080")

}
