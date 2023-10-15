package schemes

import "time"

type Menu struct {
	ID         string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	// Input with Lowercase
	Name    string `json:"name" validate:"required,lowercase,max=200" example:"master"`
	Active  *bool  `json:"active" validate:"boolean" example:"true"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	Sort    string `json:"sort"`
}

type GetMenu struct {
	ID           string    `json:"id"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	Name         string    `json:"name"`
	Active       *bool     `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
}

type MenuRequest struct {
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	// Input with Lowercase
	Name   string `json:"name" validate:"required,lowercase,max=200" example:"master"`
	Active *bool  `json:"active" validate:"boolean" example:"true"`
}
