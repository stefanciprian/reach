package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

// Get InputSource godoc
// @Summary get input-source endpoint
// @Schemes
// @Description get input-source endpoint
// @Tags input-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /input-source [get]
func GetInputSource(c *gin.Context) {
	var inputSource []models.InputSourceModel
	err := repositories.GetAllInputSources(&inputSource)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, inputSource)
	}
}

// Post InputSource godoc
// @Summary post input-source endpoint
// @Schemes
// @Description post input-source endpoint
// @Tags input-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /input-source [post]
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

// Get InputSource by Id godoc
// @Summary get input-source by id endpoint
// @Schemes
// @Description get input-source by id endpoint
// @Tags input-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /input-source/:id [get]
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

// Update InputSource godoc
// @Summary update input-source endpoint
// @Schemes
// @Description update input-source endpoint
// @Tags input-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /input-source [put]
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

// Delete InputSource godoc
// @Summary delete input-source endpoint
// @Schemes
// @Description delete input-source endpoint
// @Tags input-source
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /input-source [delete]
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
