package follow

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	service FollowServiceInterface
}

func NewFollowHand(serv FollowServiceInterface) *FollowHandler {
	return &FollowHandler{service: serv}
}

func (h *FollowHandler) FollowHandler(c *gin.Context) {

	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	followedUserID := c.Param("id")
	followedUserIDint, err := strconv.Atoi(followedUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.FollowService(ctx, userID.(int), followedUserIDint)
	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "successfully followed the user",
	})

}

func (h *FollowHandler) UnfollowHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	followedUserID := c.Param("id")
	followedUserIDint, err := strconv.Atoi(followedUserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = h.service.UnFollowService(ctx, userID.(int), followedUserIDint)
	if err != nil {
		c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "successfully unfollowed the user",
	})
}

func (h *FollowHandler) GetFollowHandler(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
	defer cancel()
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	following, err := h.service.GetFollowService(ctx, userID.(int))
	if err != nil {
		c.Error(err)

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"following": following,
	})
}
