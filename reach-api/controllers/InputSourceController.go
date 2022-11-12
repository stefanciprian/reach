package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

func GetInputSource(c *gin.Context) {
	var inputSource []models.InputSourceModel
	err := repositories.GetAllInputSources(&inputSource)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, inputSource)
	}
}

func CreateInputSource(c *gin.Context) {
	var inputSource models.InputSourceModel
	c.BindJSON(&inputSource)
	err := repositories.CreateInputSource(&inputSource)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, inputSource)
	}
}

func GetInputSourceByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var inputSource models.InputSourceModel
	err := repositories.GetInputSourceByID(&inputSource, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, inputSource)
	}
}

func UpdateInputSource(c *gin.Context) {
	var inputSource models.InputSourceModel
	id := c.Params.ByName("id")
	err := repositories.GetInputSourceByID(&inputSource, id)
	if err != nil {
		c.JSON(http.StatusNotFound, inputSource)
	}
	c.BindJSON(&inputSource)
	err = repositories.UpdateInputSource(&inputSource, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, inputSource)
	}
}

func DeleteInputSource(c *gin.Context) {
	var inputSource models.InputSourceModel
	id := c.Params.ByName("id")
	err := repositories.DeleteInputSource(&inputSource, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
