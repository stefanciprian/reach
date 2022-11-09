package controllers

import (
	"net/http"

	"reach/reach-api/services"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	msg := services.HelloService()
	c.JSON(http.StatusOK, gin.H{"status": msg})
}
