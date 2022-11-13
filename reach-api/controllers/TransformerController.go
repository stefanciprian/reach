package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

// Get Transformer godoc
// @Summary get transformer endpoint
// @Schemes
// @Description get transformer endpoint
// @Tags transformer
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /transformer [get]
func GetTransformer(c *gin.Context) {
	var transformer []models.TransformerModel
	err := repositories.GetAllTransformers(&transformer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, transformer)
	}
}

// Post Transformer godoc
// @Summary post transformer endpoint
// @Schemes
// @Description post transformer endpoint
// @Tags transformer
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /transformer [post]
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

// Get Transformer by Id godoc
// @Summary get transformer by id endpoint
// @Schemes
// @Description get transformer by id endpoint
// @Tags transformer
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /transformer/:id [get]
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

// Update Transformer godoc
// @Summary update transformer endpoint
// @Schemes
// @Description update transformer endpoint
// @Tags transformer
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /transformer [put]
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

// Delete Transformer godoc
// @Summary delete transformer endpoint
// @Schemes
// @Description delete transformer endpoint
// @Tags transformer
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /transformer [delete]
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
