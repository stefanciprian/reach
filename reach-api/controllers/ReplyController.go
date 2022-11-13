package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

// Get Reply godoc
// @Summary get reply endpoint
// @Schemes
// @Description get reply endpoint
// @Tags reply
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /reply [get]
func GetReply(c *gin.Context) {
	var reply []models.ReplyModel
	err := repositories.GetAllReplies(&reply)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reply)
	}
}

// Post Reply godoc
// @Summary post reply endpoint
// @Schemes
// @Description post reply endpoint
// @Tags reply
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /reply [post]
func CreateReply(c *gin.Context) {
	var reply models.ReplyModel
	c.BindJSON(&reply)
	err := repositories.CreateReply(&reply)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reply)
	}
}

// Get Reply by Id godoc
// @Summary get reply by id endpoint
// @Schemes
// @Description get reply by id endpoint
// @Tags reply
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /reply/:id [get]
func GetReplyByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var reply models.ReplyModel
	err := repositories.GetReplyByID(&reply, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reply)
	}
}

// Update Reply godoc
// @Summary update reply endpoint
// @Schemes
// @Description update reply endpoint
// @Tags reply
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /reply [put]
func UpdateReply(c *gin.Context) {
	var reply models.ReplyModel
	id := c.Params.ByName("id")
	err := repositories.GetReplyByID(&reply, id)
	if err != nil {
		c.JSON(http.StatusNotFound, reply)
	}
	c.BindJSON(&reply)
	err = repositories.UpdateReply(&reply, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reply)
	}
}

// Delete Reply godoc
// @Summary delete reply endpoint
// @Schemes
// @Description delete reply endpoint
// @Tags reply
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /reply [delete]
func DeleteReply(c *gin.Context) {
	var reply models.ReplyModel
	id := c.Params.ByName("id")
	err := repositories.DeleteReply(&reply, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
