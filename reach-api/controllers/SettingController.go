package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stefanciprian/reach/reach-api/models"
	"github.com/stefanciprian/reach/reach-api/repositories"
)

func GetSettings(c *gin.Context) {
	var setting []models.SettingModel
	err := repositories.GetAllSettings(&setting)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, setting)
	}
}

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
