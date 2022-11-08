package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stefanciprian/reach/reach-api/models"
)

//GetUserTests ... Get all userTests
func GetUserTests(c *gin.Context) {
	var userTest []models.UserTest
	err := models.GetAllUserTests(&userTest)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userTest)
	}
}

//CreateUserTest ... Create UserTest
func CreateUserTest(c *gin.Context) {
	var userTest models.UserTest
	c.BindJSON(&userTest)
	err := models.CreateUserTest(&userTest)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userTest)
	}
}

//GetUserTestByID ... Get the userTest by id
func GetUserTestByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var userTest models.UserTest
	err := models.GetUserTestByID(&userTest, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userTest)
	}
}

//UpdateUserTest ... Update the userTest information
func UpdateUserTest(c *gin.Context) {
	var userTest models.UserTest
	id := c.Params.ByName("id")
	err := models.GetUserTestByID(&userTest, id)
	if err != nil {
		c.JSON(http.StatusNotFound, userTest)
	}
	c.BindJSON(&userTest)
	err = models.UpdateUserTest(&userTest, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, userTest)
	}
}

//DeleteUserTest ... Delete the userTest
func DeleteUserTest(c *gin.Context) {
	var userTest models.UserTest
	id := c.Params.ByName("id")
	err := models.DeleteUserTest(&userTest, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
