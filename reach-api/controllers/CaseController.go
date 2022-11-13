package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

// Get Case godoc
// @Summary get case endpoint
// @Schemes
// @Description get case endpoint
// @Tags case
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /case [get]
func GetCase(c *gin.Context) {
	var caseModel []models.CaseModel
	err := repositories.GetAllCases(&caseModel)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, caseModel)
	}
}

// Post Case godoc
// @Summary post case endpoint
// @Schemes
// @Description post case endpoint
// @Tags case
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /case [post]
func CreateCase(c *gin.Context) {
	var caseModel models.CaseModel
	c.BindJSON(&caseModel)
	err := repositories.CreateCase(&caseModel)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, caseModel)
	}
}

// Get Case by Id godoc
// @Summary get case by id endpoint
// @Schemes
// @Description get case by id endpoint
// @Tags case
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /case/:id [get]
func GetCaseByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var caseModel models.CaseModel
	err := repositories.GetCaseByID(&caseModel, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, caseModel)
	}
}

// Update Case godoc
// @Summary update case endpoint
// @Schemes
// @Description update case endpoint
// @Tags case
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /case [put]
func UpdateCase(c *gin.Context) {
	var caseModel models.CaseModel
	id := c.Params.ByName("id")
	err := repositories.GetCaseByID(&caseModel, id)
	if err != nil {
		c.JSON(http.StatusNotFound, caseModel)
	}
	c.BindJSON(&caseModel)
	err = repositories.UpdateCase(&caseModel, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, caseModel)
	}
}

// Delete Case godoc
// @Summary delete case endpoint
// @Schemes
// @Description delete case endpoint
// @Tags case
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /case [delete]
func DeleteCase(c *gin.Context) {
	var caseModel models.CaseModel
	id := c.Params.ByName("id")
	err := repositories.DeleteCase(&caseModel, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
