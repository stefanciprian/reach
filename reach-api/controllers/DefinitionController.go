package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

// Get Definition godoc
// @Summary get definition endpoint
// @Schemes
// @Description get definition endpoint
// @Tags definition
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /definition [get]
func GetDefinition(c *gin.Context) {
	var definition []models.DefinitionModel
	err := repositories.GetAllDefinitions(&definition)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, definition)
	}
}

// Post Definition godoc
// @Summary post definition endpoint
// @Schemes
// @Description post definition endpoint
// @Tags definition
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /definition [post]
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

// Get Definition by Id godoc
// @Summary get definition by id endpoint
// @Schemes
// @Description get definition by id endpoint
// @Tags definition
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /definition/:id [get]
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

// Update Definition godoc
// @Summary update definition endpoint
// @Schemes
// @Description update definition endpoint
// @Tags definition
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /definition [put]
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

// Delete Definition godoc
// @Summary delete definition endpoint
// @Schemes
// @Description delete definition endpoint
// @Tags definition
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /definition [delete]
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
