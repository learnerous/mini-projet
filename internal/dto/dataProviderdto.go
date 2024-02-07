package dto

import "time"

type Get struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Country   string `json:"country"`
	WebSite   string `json:"website"`
	DataTypes string `json:"data_types"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
type Update struct {
	Name      string    `form:"name" json:"name"`
	Email     string    `form:"email" json:"email"`
	Country   string    `form:"country" json:"country"`
	WebSite   string    `form:"website" json:"website"`
	DataTypes string    `form:"dataTypes" json:"data_types"`
	CreatedAt time.Time `form:"createdAt" json:"created_at"`
	UpdatedAt time.Time `form:"updatedAt" json:"updated_at"`
}
type Create struct {
	Update
}
