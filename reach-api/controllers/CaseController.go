package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

func GetCase(c *gin.Context) {
	var caseModel []models.CaseModel
	err := repositories.GetAllCases(&caseModel)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, caseModel)
	}
}

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
