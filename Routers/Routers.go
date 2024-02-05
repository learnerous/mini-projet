package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/path/to/controllers" // Import the "controllers" package
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	dp := r.Group("/data-providers")
	{
		dp.GET("/", controllers.HelloDataProvider) // Use the correct syntax to reference the "HelloDataProvider" function
	}
}
