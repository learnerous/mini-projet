package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//create a new data provider
// func (c *controller) CreateDataProvider(ctx *gin.Context) {
// 	var dataProvider models.DataProvider
// 	err := ctx.BindJSON(&dataProvider)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	err = c.dataProviderService.CreateDataProvider(&dataProvider)
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusCreated, dataProvider)
// }
var db *mongo.Collection

func GetDataProvider(c *gin.Context) {

	// if err := db.Find(&posts).All(nil); err != nil {
	//  c.AbortWithStatus(http.StatusInternalServerError)
	//  return
	// }
	c.JSON(http.StatusOK, "Hello Data Provider")
}
func CreateDataProvider(c *gin.Context) {
	fmt.Println("trying to create a new data provider", c.Params)
	c.JSON(http.StatusOK, "trying to create a new data provider")
}
