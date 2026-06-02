package post

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service PostServiceInterface
}

func NewPostHandler(s PostServiceInterface) *PostHandler {
	return &PostHandler{
		service: s,
	}
}

func (s *PostHandler) CreatePostHandler(c *gin.Context) {

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

	var userRequest CreatePost
	err := c.ShouldBindJSON(&userRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	proofReturn, err := s.service.CreatePostService(userIDInt, userRequest)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"post created": proofReturn,
	})
}

func (s *PostHandler) GetPostHandler(c *gin.Context) {
	postID := c.Param("id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	userPost, err := s.service.GetPostService(postIDint)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": userPost,
	})

}

func (s *PostHandler) UpdatePostHandler(c *gin.Context) {
	postID := c.Param("id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

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

	var UPost UpdatePost
	err = c.ShouldBindJSON(&UPost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	updatedPost, err := s.service.UpdatePostService(userIDInt, postIDint, UPost)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated_post": updatedPost,
	})

}

func (s *PostHandler) DeletePostHandler(c *gin.Context) {
	postID := c.Param("id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input",
		})
		return
	}

	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}
	userIdInt, ok := userID.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	err = s.service.DeletePostService(userIdInt, postIDint)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": "deleted",
	})

}
