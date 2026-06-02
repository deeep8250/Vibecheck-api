package follow

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	service FollowServiceInterface
}

func NewFollowHand(serv FollowServiceInterface) *FollowHandler {
	return &FollowHandler{service: serv}
}

func (h *FollowHandler) FollowHandler(c *gin.Context) {

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

	err = h.service.FollowService(userID.(int), followedUserIDint)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "successfully followed the user",
	})

}

func (h *FollowHandler) UnfollowHandler(c *gin.Context) {

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

	err = h.service.UnFollowService(userID.(int), followedUserIDint)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "successfully unfollowed the user",
	})
}

func (h *FollowHandler) GetFollowHandler(c *gin.Context) {

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	following, err := h.service.GetFollowService(userID.(int))
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"following": following,
	})
}
