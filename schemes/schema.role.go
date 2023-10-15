package schemes

import "time"

type Role struct {
	ID string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name string `json:"name" validate:"required,lowercase,max=200" example:"root"`
	// Input with Lowercase
	Type       string `json:"type" validate:"required,lowercase,max=100" example:"sys"`
	MerchantID string `json:"merchant_id" validate:"required,uuid"`
	Active     *bool  `json:"active" validate:"boolean" example:"true"`
	Page       int    `json:"page"`
	PerPage    int    `json:"perpage"`
	Sort       string `json:"sort"`
}

type GetRole struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type GetAllRole struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	Active       *bool     `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
}

type RoleRequest struct {
	// Input with Lowercase
	Name string `json:"name" validate:"required,lowercase,max=200" example:"root"`
	// Input with Lowercase
	Type       string `json:"type" validate:"required,lowercase,max=100" example:"sys"`
	MerchantID string `json:"merchant_id" validate:"required,uuid" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	Active     *bool  `json:"active" validate:"boolean" example:"true"`
}
