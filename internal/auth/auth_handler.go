package auth

import (
	"net/http"

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

	User, err := h.service.Register(userInput)
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

	token, err := h.service.Login(userInput)
	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
