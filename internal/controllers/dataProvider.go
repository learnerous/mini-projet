package controllers

import (
	"mini_projet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//create a new data provider
func (c *Controller) CreateDataProvider(ctx *gin.Context) {
	var dataProvider models.DataProvider
	err := ctx.BindJSON(&dataProvider)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = c.dataProviderService.CreateDataProvider(&dataProvider)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, dataProvider)
}
func helloDataProvider(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "Hello DataProvider"})
}
