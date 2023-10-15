package schemes

import "time"

type MenuDetailFunction struct {
	ID           string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID   string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	MenuID       string `json:"menu_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuDetailID string `json:"menu_detail_id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name    string `json:"name" validate:"required,lowercase,max=200" example:"add produk"`
	Link    string `json:"link"  validate:"required" example:"/add"`
	Active  *bool  `json:"active" validate:"boolean" example:"true"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	Sort    string `json:"sort"`
}

type GetMenuDetailFunction struct {
	ID             string    `json:"id"`
	MerchantID     string    `json:"merchant_id"`
	MerchantName   string    `json:"merchant_name"`
	MenuID         string    `json:"menu_id"`
	MenuName       string    `json:"menu_name"`
	MenuDetailID   string    `json:"menu_detail_id"`
	MenuDetailName string    `json:"menu_detail_name"`
	Name           string    `json:"name"`
	Link           string    `json:"link"`
	Active         *bool     `json:"active"`
	CreatedAt      time.Time `json:"created_at"`
}

type MenuDetailFunctionRequest struct {
	MerchantID   string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	MenuID       string `json:"menu_id" validate:"required,uuid" example:"890e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MenuDetailID string `json:"menu_detail_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name   string `json:"name" validate:"required,lowercase,max=200" example:"add produk"`
	Link   string `json:"link" validate:"required" example:"/add"`
	Active *bool  `json:"active" validate:"boolean" example:"true"`
}
