package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stefanciprian/reach/services"
)

func Hello(c *gin.Context) {
	msg := services.HelloService()
	c.JSON(http.StatusOK, gin.H{"status": msg})
}
