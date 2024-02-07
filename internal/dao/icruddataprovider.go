package dao

import (
	"context"
	"mini_projet/models"
)

type ICRUDDataProvider interface {
	Create(ctx context.Context, dataProvider *models.DataProvider) (*models.DataProvider, error)
	GetDataProvider(ctx context.Context) ([]*models.DataProvider, error)
	GetDataProviderByID(ctx context.Context, id string) (*models.DataProvider, error)
	UpdateDataProvider(context.Context, string, *models.DataProvider) (*models.DataProvider, error)
	DeleteDataProvider(ctx context.Context, id string) (bool, error)
}
