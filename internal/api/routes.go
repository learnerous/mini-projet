package api

import "github.com/gin-gonic/gin"

func initRoutes(router gin.IRouter, controller *controller.Controller) {
	router.GET("/data-providers", controller.GetDataProvider)
	router.GET("/data-providers/:id", controller.GetDataProviderByID)
	router.POST("/data-providers", controller.CreateDataProvider)
	router.PUT("/data-providers/:id", controller.UpdateDataProvider)
	router.DELETE("/data-providers/:id", controller.DeleteDataProvider)
}
