package controllers

import (
	"fmt"
	"net/http"

	"reach/reach-api/models"
	"reach/reach-api/repositories"

	"github.com/gin-gonic/gin"
)

// Get User godoc
// @Summary get user endpoint
// @Schemes
// @Description get user endpoint
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /users/user [get]
func GetUser(c *gin.Context) {
	var user []models.UserModel
	err := repositories.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Post User godoc
// @Summary post user endpoint
// @Schemes
// @Description post user endpoint
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /users/user [post]
func CreateUser(c *gin.Context) {
	var user models.UserModel
	c.BindJSON(&user)
	err := repositories.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Get User by Id godoc
// @Summary get user by id endpoint
// @Schemes
// @Description get user by id endpoint
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /users/user/:id [get]
func GetUserByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.UserModel
	err := repositories.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Update User godoc
// @Summary update user endpoint
// @Schemes
// @Description update user endpoint
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /users/user [put]
func UpdateUser(c *gin.Context) {
	var user models.UserModel
	id := c.Params.ByName("id")
	err := repositories.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = repositories.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

// Delete User godoc
// @Summary delete user endpoint
// @Schemes
// @Description delete user endpoint
// @Tags user
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /users/user [delete]
func DeleteUser(c *gin.Context) {
	var user models.UserModel
	id := c.Params.ByName("id")
	err := repositories.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}
