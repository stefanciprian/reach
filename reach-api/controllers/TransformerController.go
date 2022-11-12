package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

func GetTransformer(c *gin.Context) {
	var transformer []models.TransformerModel
	err := repositories.GetAllTransformers(&transformer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transformer)
	}
}

func CreateTransformer(c *gin.Context) {
	var transformer models.TransformerModel
	c.BindJSON(&transformer)
	err := repositories.CreateTransformer(&transformer)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transformer)
	}
}

func GetTransformerByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var transformer models.TransformerModel
	err := repositories.GetTransformerByID(&transformer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transformer)
	}
}

func UpdateTransformer(c *gin.Context) {
	var transformer models.TransformerModel
	id := c.Params.ByName("id")
	err := repositories.GetTransformerByID(&transformer, id)
	if err != nil {
		c.JSON(http.StatusNotFound, transformer)
	}
	c.BindJSON(&transformer)
	err = repositories.UpdateTransformer(&transformer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transformer)
	}
}

func DeleteTransformer(c *gin.Context) {
	var transformer models.TransformerModel
	id := c.Params.ByName("id")
	err := repositories.DeleteTransformer(&transformer, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
