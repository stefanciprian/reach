package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

// Get Setting godoc
// @Summary get setting endpoint
// @Schemes
// @Description get setting endpoint
// @Tags setting
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /settings/setting [get]
func GetSettings(c *gin.Context) {
	var setting []models.SettingModel
	err := repositories.GetAllSettings(&setting)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

// Post Setting godoc
// @Summary post setting endpoint
// @Schemes
// @Description post setting endpoint
// @Tags setting
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /settings/setting [post]
func CreateSetting(c *gin.Context) {
	var setting models.SettingModel
	c.BindJSON(&setting)
	err := repositories.CreateSetting(&setting)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

// Get Setting by Id godoc
// @Summary get setting by id endpoint
// @Schemes
// @Description get setting by id endpoint
// @Tags setting
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /settings/setting/:id [get]
func GetSettingByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var setting models.SettingModel
	err := repositories.GetSettingByID(&setting, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

// Update Setting godoc
// @Summary update setting endpoint
// @Schemes
// @Description update setting endpoint
// @Tags setting
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /settings/setting [put]
func UpdateSetting(c *gin.Context) {
	var setting models.SettingModel
	id := c.Params.ByName("id")
	err := repositories.GetSettingByID(&setting, id)
	if err != nil {
		c.JSON(http.StatusNotFound, setting)
	}
	c.BindJSON(&setting)
	err = repositories.UpdateSetting(&setting, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

// Delete Setting godoc
// @Summary delete setting endpoint
// @Schemes
// @Description delete setting endpoint
// @Tags setting
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /settings/setting [delete]
func DeleteSetting(c *gin.Context) {
	var setting models.SettingModel
	id := c.Params.ByName("id")
	err := repositories.DeleteSetting(&setting, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
