package routes

import (
	"net/http"

	"github.com/stefanciprian/reach/controllers"

	"github.com/stefanciprian/reach/config"

	"github.com/bamzi/jobrunner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	sess := config.ConnectAws()

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("sess", sess)
		c.Next()
	})
	router.Use(gin.Logger())
	router.Use(cors.Default())

	router.LoadHTMLGlob("templates/*.tmpl.html")

	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	// Resource to return the JSON data
	router.GET("/jobrunner/json", func(c *gin.Context) {
		// returns a map[string]interface{} that can be marshalled as JSON
		c.JSON(http.StatusOK, jobrunner.StatusJson())
	})

	// Returns html page at given endpoint based on the loaded
	router.GET("/jobrunner/html", func(c *gin.Context) {
		// Returns the template data pre-parsed
		c.HTML(http.StatusOK, "status.tmpl.html", jobrunner.StatusPage())
	})

	router.GET("hello", controllers.Hello)

	return router
}
