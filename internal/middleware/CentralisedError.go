package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandle() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last()
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Err.Error(),
			})
			return
		}

	}

}
