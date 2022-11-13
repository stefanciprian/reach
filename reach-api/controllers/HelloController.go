package controllers

import (
	"net/http"

	"reach/reach-api/services"

	"github.com/gin-gonic/gin"
)

type HelloResponse struct {
	status string
}

// Hello godoc
// @Summary hello endpoint
// @Schemes
// @Description say hello
// @Tags hello
// @Accept json
// @Produce json
// @Success 200 {param HelloResponse} HelloResponse
// @Router /hello [get]
func Hello(c *gin.Context) {
	msg := services.HelloService()
	c.JSON(http.StatusOK, gin.H{"status": msg})
}
