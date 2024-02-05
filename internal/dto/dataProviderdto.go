package dto

type CreateDataProvider struct {
	OID       int    `json:"_id"`
	Name      string `json:"firstname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	WebSite   string `json:"website"`
	DataTypes string `json:"data_types"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
