package internal

import (
	"mini_projet/internal/api/controllers"
	"mini_projet/internal/dao/mongodbimplementation"
	"mini_projet/pkg/validation"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	controller := controllers.Controller{
		CRUDDataProvider: &mongodbimplementation.MongoDBCRUDDataProvider{},
		Validate:         validation.DtoValidator(),
	}

	dp := r.Group("/data-providers")
	{
		//dp.GET("/", controllers.GetDataProvider)
		dp.POST("/add", controller.CreateDataProvider)
		dp.GET("/get/:id", controller.GetDataProvider)
		dp.GET("/get", controller.GetAllDataProviders)
		dp.DELETE("/delete/:id", controller.DeleteDataProvider)
		dp.PUT("/update/:id", controller.UpdateDataProvider)
	}
	return r
}
