package schemes

import "time"

type ProductCategory struct {
	ID         string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID string `json:"merchant_id" validate:"required,uuid" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	OutletID   string `json:"outlet_id" validate:"required,uuid" example:"4e769a02-0214-4277-90d0-bdf7f7b7c064"`
	// Input with Lowercase
	Name    string `json:"name" validate:"required,lowercase,max=200" example:"mie"`
	Active  *bool  `json:"active" validate:"boolean" example:"true"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	Sort    string `json:"sort"`
}

type GetProductCategory struct {
	ID           string    `json:"id"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	OutletID     string    `json:"outlet_id"`
	OutletName   string    `json:"outlet_name"`
	Name         string    `json:"name"`
	Active       *bool     `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
}

type ProductCategoryRequest struct {
	MerchantID string `json:"merchant_id" validate:"required,uuid" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	OutletID   string `json:"outlet_id" validate:"required,uuid" example:"4e769a02-0214-4277-90d0-bdf7f7b7c064"`
	// Input with Lowercase
	Name   string `json:"name" validate:"required,lowercase,max=200" example:"mie"`
	Active *bool  `json:"active" validate:"boolean" example:"true"`
}
