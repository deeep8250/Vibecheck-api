package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service AuthServiceInterface
}

func NewAuthHandler(serv AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		service: serv,
	}
}

func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var userInput Register
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	User, err := h.service.Register(ctx, userInput)
	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"User": User,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var userInput Login
	err := c.ShouldBindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	token, err := h.service.Login(ctx, userInput)
	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
