package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

func GetReply(c *gin.Context) {
	var reply []models.ReplyModel
	err := repositories.GetAllReplies(&reply)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, reply)
	}
}

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
