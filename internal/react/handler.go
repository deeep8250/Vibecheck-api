package react

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReactHandler struct {
	srv ReactServiceInterface
}

func NewReactHandler(s ReactServiceInterface) *ReactHandler {
	return &ReactHandler{
		srv: s,
	}
}

func (h *ReactHandler) ReactPost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	postID := c.Param("id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	type EmojiInput struct {
		emoji string `json:"emoji"`
	}
	var inp EmojiInput
	err = c.ShouldBindJSON(&inp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.srv.ReactPostServ(ctx, postIDint, userIDInt, inp.emoji)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"response": "reacted the post",
	})

}

func (h *ReactHandler) GetReactionByPost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	postID := c.Param("id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	result, err := h.srv.GetReactionsByPostID(ctx, postIDint)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": result,
	})

}

func (h *ReactHandler) DeleteReaction(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	userIDInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	postID := c.Param("id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	err = h.srv.RemoveReact(ctx, postIDint, userIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"response": "reaction deleted",
	})

}
