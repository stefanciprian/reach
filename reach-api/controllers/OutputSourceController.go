package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

// Get OutputSource godoc
// @Summary get output-source endpoint
// @Schemes
// @Description get output-source endpoint
// @Tags output-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /output-source [get]
func GetOutputSource(c *gin.Context) {
	var outputSource []models.OutputSourceModel
	err := repositories.GetAllOutputSources(&outputSource)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, outputSource)
	}
}

// Post OutputSource godoc
// @Summary post output-source endpoint
// @Schemes
// @Description post output-source endpoint
// @Tags output-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /output-source [post]
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

// Get OutputSource by Id godoc
// @Summary get output-source by id endpoint
// @Schemes
// @Description get output-source by id endpoint
// @Tags output-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /output-source/:id [get]
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

// Update OutputSource godoc
// @Summary update output-source endpoint
// @Schemes
// @Description update output-source endpoint
// @Tags output-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /output-source [put]
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

// Delete OutputSource godoc
// @Summary delete output-source endpoint
// @Schemes
// @Description delete output-source endpoint
// @Tags output-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /output-source [delete]
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
