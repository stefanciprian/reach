package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

func GetOutputSource(c *gin.Context) {
	var outputSource []models.OutputSourceModel
	err := repositories.GetAllOutputSources(&outputSource)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, outputSource)
	}
}

func CreateOutputSource(c *gin.Context) {
	var outputSource models.OutputSourceModel
	c.BindJSON(&outputSource)
	err := repositories.CreateOutputSource(&outputSource)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, outputSource)
	}
}

func GetOutputSourceByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var outputSource models.OutputSourceModel
	err := repositories.GetOutputSourceByID(&outputSource, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, outputSource)
	}
}

func UpdateOutputSource(c *gin.Context) {
	var outputSource models.OutputSourceModel
	id := c.Params.ByName("id")
	err := repositories.GetOutputSourceByID(&outputSource, id)
	if err != nil {
		c.JSON(http.StatusNotFound, outputSource)
	}
	c.BindJSON(&outputSource)
	err = repositories.UpdateOutputSource(&outputSource, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, outputSource)
	}
}

func DeleteOutputSource(c *gin.Context) {
	var outputSource models.OutputSourceModel
	id := c.Params.ByName("id")
	err := repositories.DeleteOutputSource(&outputSource, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
