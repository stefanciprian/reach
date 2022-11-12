package routes

import (
	"net/http"

	"reach/reach-api/config"
	"reach/reach-api/controllers"

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

		api.GET("case", controllers.GetCase)
		api.POST("case", controllers.CreateCase)
		api.GET("case/:id", controllers.GetCaseByID)
		api.PUT("case/:id", controllers.UpdateCase)
		api.DELETE("case/:id", controllers.DeleteCase)

		api.GET("definition", controllers.GetDefinition)
		api.POST("definition", controllers.CreateDefinition)
		api.GET("definition/:id", controllers.GetDefinitionByID)
		api.PUT("definition/:id", controllers.UpdateDefinition)
		api.DELETE("definition/:id", controllers.DeleteDefinition)

		api.GET("input-source", controllers.GetInputSource)
		api.POST("input-source", controllers.CreateInputSource)
		api.GET("input-source/:id", controllers.GetInputSourceByID)
		api.PUT("input-source/:id", controllers.UpdateInputSource)
		api.DELETE("input-source/:id", controllers.DeleteInputSource)

		api.GET("output-source", controllers.GetOutputSource)
		api.POST("output-source", controllers.CreateOutputSource)
		api.GET("output-source/:id", controllers.GetOutputSourceByID)
		api.PUT("output-source/:id", controllers.UpdateOutputSource)
		api.DELETE("output-source/:id", controllers.DeleteOutputSource)

		api.GET("reply", controllers.GetReply)
		api.POST("reply", controllers.CreateReply)
		api.GET("reply/:id", controllers.GetReplyByID)
		api.PUT("reply/:id", controllers.UpdateReply)
		api.DELETE("reply/:id", controllers.DeleteReply)

		api.GET("transformer", controllers.GetTransformer)
		api.POST("transformer", controllers.CreateTransformer)
		api.GET("transformer/:id", controllers.GetTransformerByID)
		api.PUT("transformer/:id", controllers.UpdateTransformer)
		api.DELETE("transformer/:id", controllers.DeleteTransformer)
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
