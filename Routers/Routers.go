package routers

import (
	"mini_projet/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	dp := r.Group("/data-providers")
	{
		dp.GET("/", controllers.GetDataProvider)
		dp.POST("/add", controllers.CreateDataProvider)
	}
	return r
}
