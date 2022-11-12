package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

func GetDefinition(c *gin.Context) {
	var definition []models.DefinitionModel
	err := repositories.GetAllDefinitions(&definition)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, definition)
	}
}

func CreateDefinition(c *gin.Context) {
	var definition models.DefinitionModel
	c.BindJSON(&definition)
	err := repositories.CreateDefinition(&definition)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, definition)
	}
}

func GetDefinitionByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var definition models.DefinitionModel
	err := repositories.GetDefinitionByID(&definition, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, definition)
	}
}

func UpdateDefinition(c *gin.Context) {
	var definition models.DefinitionModel
	id := c.Params.ByName("id")
	err := repositories.GetDefinitionByID(&definition, id)
	if err != nil {
		c.JSON(http.StatusNotFound, definition)
	}
	c.BindJSON(&definition)
	err = repositories.UpdateDefinition(&definition, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, definition)
	}
}

func DeleteDefinition(c *gin.Context) {
	var definition models.DefinitionModel
	id := c.Params.ByName("id")
	err := repositories.DeleteDefinition(&definition, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
