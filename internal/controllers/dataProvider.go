package controllers

import (
	"fmt"
	"mini_projet/internal/dao"
	"mini_projet/internal/dto"
	"net/http"

	"bitbucket.org/amaltheafs/pkg/logutil"
	"bitbucket.org/amaltheafs/tenant-server/pkg/serializers"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	CRUDDataProvider dao.ICRUDDataProvider
}

func (Controller *Controller) CreateDataProvider(c *gin.Context) {

	var dto dto.CreateDataProvider
	err := c.ShouldBind(&dto)
	fmt.Println("dto", dto)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}
	dataprovider, err := serializers.DeserializeCreateDto(dto)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, "Invalid request")
		return
	}
	dataprovider, err = Controller.CRUDDataProvider.CreateDataProvider(c, dataprovider)
	getDto, err := serializers.Serialize(*dataprovider)
	c.JSON(http.StatusOK, getDto)
	//raw, err := c.GetRawData()
	// var dataProviderdto models.DataProvider

	// body, err := ioutil.ReadAll(c.Request.Body)
	// err = c.BindJSON(dataProviderdto)

	// fmt.Println("trying to create a new data provider with id :", string(body), dataProviderdto, err)
	// c.JSON(http.StatusOK, "trying to create a new data provider")
}
func GetDataProvider(c *gin.Context) {

	// if err := db.Find(&posts).All(nil); err != nil {
	//  c.AbortWithStatus(http.StatusInternalServerError)
	//  return
	// }
	c.JSON(http.StatusOK, "Hello Data Provider")
}

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
