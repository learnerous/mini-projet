package routers

import (
	"github.com/gin-gonic/gin"
)

// func helloDataProvider(c *gin.Context) {
// 	// Implement the logic for helloDataProvider here
// }

func SetupRouter() *gin.Engine {
	r := gin.Default()
	dp := r.Group("/data-providers")
	{
		dp.GET("/", controllers.helloDataProvider)
	}
	return r
}
