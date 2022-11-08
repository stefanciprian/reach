package routes

import (
	"net/http"

	"github.com/stefanciprian/reach/reach-api/controllers"

	"github.com/stefanciprian/reach/reach-api/config"

	"github.com/bamzi/jobrunner"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
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

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile("./reach-client/build", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("hello", controllers.Hello)

		// Resource to return the JSON data
		api.GET("/jobrunner/json", func(c *gin.Context) {
			// returns a map[string]interface{} that can be marshalled as JSON
			c.JSON(http.StatusOK, jobrunner.StatusJson())
		})
	}

	userTests := api.Group("/users")
	{
		userTests.GET("user", controllers.GetUser)
		userTests.POST("user", controllers.CreateUser)
		userTests.GET("user/:id", controllers.GetUserByID)
		userTests.PUT("user/:id", controllers.UpdateUser)
		userTests.DELETE("user/:id", controllers.DeleteUser)
	}
	settings := api.Group("/settings")
	{
		settings.GET("setting", controllers.GetSettings)
		settings.POST("setting", controllers.CreateSetting)
		settings.GET("setting/:id", controllers.GetSettingByID)
		settings.PUT("setting/:id", controllers.UpdateSetting)
		settings.DELETE("setting/:id", controllers.DeleteSetting)
	}

	return router
}
