package serializers

import (
	"mini_projet/models"
	"time"

	"mini_projet/internal/dto"
)

func DeserializeCreateDto(dtoContent dto.Create /*, id primitive.ObjectID*/) (*models.DataProvider, error) {

	createdAt := time.Now()
	updatedAt := time.Now()
	DP := &models.DataProvider{
		Name:      dtoContent.Name,
		Email:     dtoContent.Email,
		Country:   dtoContent.Country,
		WebSite:   dtoContent.WebSite,
		DataTypes: dtoContent.DataTypes,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	return DP, nil
}
func Serialize(dataProvifer models.DataProvider) (*dto.Get, error) {
	dto := &dto.Get{
		ID:        dataProvifer.OID.Hex(),
		Name:      dataProvifer.Name,
		Email:     dataProvifer.Email,
		Country:   dataProvifer.Country,
		WebSite:   dataProvifer.WebSite,
		DataTypes: dataProvifer.DataTypes,
		CreatedAt: dataProvifer.CreatedAt.String(), // Convert time.Time to string
		UpdatedAt: dataProvifer.UpdatedAt.String(),
	}

	return dto, nil
}
func DeserializeUpdateDto(dtoUpdate dto.Update) (*models.DataProvider, error) {
	dtoCreate := dto.Create{
		Update: dtoUpdate,
	}
	return DeserializeCreateDto(dtoCreate)

}
