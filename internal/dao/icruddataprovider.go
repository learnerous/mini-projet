package dao

import "mini_projet/models"

type ICRUDDataProvider interface {
	CreateDataProvider(dataProvider *models.DataProvider) (*models.DataProvider, error)
	GetDataProvider() ([]models.DataProvider, error)
	GetDataProviderByID(id string) (*models.DataProvider, error)
	UpdateDataProvider(dataProvider *models.DataProvider) (*models.DataProvider, error)
	DeleteDataProvider(id string) error
}
