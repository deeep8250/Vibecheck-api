package feed

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FeedHandler struct {
	service FeedServiceInterface
}

func NewFeedHandler(s FeedServiceInterface) *FeedHandler {
	return &FeedHandler{
		service: s,
	}
}

func (h *FeedHandler) GetFeed(c *gin.Context) {
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

	page := c.DefaultQuery("page", "1")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page is empty",
		})
		return
	}
	if pageInt < 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "page must be greater than 0",
		})
		return
	}

	limit := c.DefaultQuery("limit", "5")
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "limit is empty",
		})
		return
	}
	if limitInt > 50 || limitInt <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "limit must be 0  to  50",
		})
		return
	}

	feed, err := h.service.GetFeed(ctx, userIDInt, limitInt, pageInt)
	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"feed": feed,
	})
}
