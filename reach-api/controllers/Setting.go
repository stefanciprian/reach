package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stefanciprian/reach/reach-api/models"
)

//GetSettings ... Get all settings
func GetSettings(c *gin.Context) {
	var setting []models.Setting
	err := models.GetAllSettings(&setting)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

//CreateSettings ... Create Settings
func CreateSetting(c *gin.Context) {
	var setting models.Setting
	c.BindJSON(&setting)
	err := models.CreateSetting(&setting)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

//GetSettingByID ... Get the setting by id
func GetSettingByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var setting models.Setting
	err := models.GetSettingByID(&setting, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

//UpdateSetting ... Update the setting information
func UpdateSetting(c *gin.Context) {
	var setting models.Setting
	id := c.Params.ByName("id")
	err := models.GetSettingByID(&setting, id)
	if err != nil {
		c.JSON(http.StatusNotFound, setting)
	}
	c.BindJSON(&setting)
	err = models.UpdateSetting(&setting, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

//DeleteSetting ... Delete the setting
func DeleteSetting(c *gin.Context) {
	var setting models.Setting
	id := c.Params.ByName("id")
	err := models.DeleteSetting(&setting, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
