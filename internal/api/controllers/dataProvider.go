package controllers

import (
	"fmt"
	"mini_projet/internal/dao"
	"mini_projet/internal/dto"
	errs "mini_projet/pkg/core/errors"
	"mini_projet/pkg/logutil"
	"mini_projet/pkg/rest"
	"mini_projet/pkg/rest/errorcodes/generic"
	dtovalidation "mini_projet/pkg/validation"
	"mini_projet/serializers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Controller struct {
	CRUDDataProvider dao.ICRUDDataProvider
	Validate         *validator.Validate
}

func (controller *Controller) CreateDataProvider(c *gin.Context) {
	response := rest.ResponseShape{Code: generic.UNKNOWN.String()}

	var dto dto.Create

	err := c.ShouldBind(&dto)

	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = controller.Validate.Struct(&dto)
	incorrectFields := dtovalidation.ParseErrors(err)
	if incorrectFields != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, rest.ResponseShape{Fields: incorrectFields})
		return
	}
	dataprovider, err := serializers.DeserializeCreateDto(dto /*, DpId*/)
	fmt.Println("dataprovider\n in controller", dataprovider, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, response)
		return
	}

	dataprovider, err = controller.CRUDDataProvider.Create(c, dataprovider)

	if err != nil {
		fmt.Println("error", err)
		logutil.Logger().Errorf("%s", err)
		if customErr, ok := err.(*errs.CustomError); ok {
			response.Code = customErr.Code
			response.Context = customErr.Type
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	getDto, err := serializers.Serialize(*dataprovider)
	c.JSON(
		http.StatusOK,
		rest.ResponseShape{
			Data: getDto,
		},
	)
}
func (controller *Controller) GetDataProvider(c *gin.Context) {
	id := c.Param("id")
	dataprovider, err := controller.CRUDDataProvider.GetDataProviderByID(c, id)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(
			http.StatusBadRequest,
			rest.ResponseShape{Code: generic.UNKNOWN.String()},
		)
		return
	}
	if dataprovider == nil {
		c.JSON(
			http.StatusNotFound,
			rest.ResponseShape{Code: generic.DB_RECORD_NOT_FOUND.String()},
		)
		return
	}
	dto, err := serializers.Serialize(*dataprovider)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(
			http.StatusBadRequest,
			rest.ResponseShape{Code: generic.UNKNOWN.String()},
		)
		return
	}

	c.JSON(http.StatusOK, rest.ResponseShape{Data: dto})
}
func (controller *Controller) GetAllDataProviders(c *gin.Context) {
	response := rest.ResponseShape{Code: generic.UNKNOWN.String()}
	dataproviders, err := controller.CRUDDataProvider.GetDataProvider(c)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, response)
		return

	}
	dtoDataproviders := make([]dto.Get, len(dataproviders))
	for i, dataprovider := range dataproviders {
		dto, err := serializers.Serialize(*dataprovider)
		if err != nil {
			logutil.Logger().Errorf("%s", err)
			c.JSON(http.StatusBadRequest, response)
			return
		}
		dtoDataproviders[i] = *dto

	}
	response.Data = dtoDataproviders
	response.Code = ""
	c.JSON(http.StatusOK, response)
}
func (controller *Controller) DeleteDataProvider(c *gin.Context) {
	id := c.Param("id")

	deleted, err := controller.CRUDDataProvider.DeleteDataProvider(c, id)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, rest.ResponseShape{Code: generic.UNKNOWN.String()})
		return
	}

	if !deleted {
		logutil.Logger().Infof("Record not found")
		c.JSON(http.StatusNotFound, rest.ResponseShape{Code: generic.DB_RECORD_NOT_FOUND.String()})
		return
	}

	c.JSON(http.StatusOK, rest.ResponseShape{Data: id})
}
func (controller *Controller) UpdateDataProvider(c *gin.Context) {
	id := c.Param("id")
	response := rest.ResponseShape{Code: generic.UNKNOWN.String()}

	var dto dto.Update
	err := c.ShouldBind(&dto)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	err = controller.Validate.Struct(&dto)
	incorrectFields := dtovalidation.ParseErrors(err)
	if incorrectFields != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, rest.ResponseShape{Fields: incorrectFields})
		return
	}
	dataprovider, err := serializers.DeserializeUpdateDto(dto)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	dataprovider, err = controller.CRUDDataProvider.UpdateDataProvider(c, id, dataprovider)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		if customErr, ok := err.(*errs.CustomError); ok {
			response.Code = customErr.Code
			response.Context = customErr.Type
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}
	getDto, err := serializers.Serialize(*dataprovider)
	if err != nil {
		logutil.Logger().Errorf("%s", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(
		http.StatusOK,
		rest.ResponseShape{
			Data: getDto,
		},
	)

}
