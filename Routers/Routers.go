package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/learnerous/mini-projet/internal/controllers"
)

// func helloDataProvider(c *gin.Context) {
// 	// Implement the logic for helloDataProvider here
// }

package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/path/to/controllers" // Import the controllers package
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	dp := r.Group("/data-providers")
	{
		dp.GET("/", controllers.helloDataProvider)
	}
	return r
}
