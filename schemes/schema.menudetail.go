package schemes

import "time"

type MenuDetail struct {
	ID         string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	MenuID     string `json:"menu_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	// Input with Lowercase
	Name    string `json:"name" validate:"required,lowercase,max=200" example:"master produk"`
	Link    string `json:"link"  validate:"required" example:"/master-produk"`
	Image   string `json:"image"`
	Icon    string `json:"icon"`
	Active  *bool  `json:"active" validate:"boolean" example:"true"`
	Page    int    `json:"page"`
	PerPage int    `json:"perpage"`
	Sort    string `json:"sort"`
}

type GetMenuDetail struct {
	ID           string    `json:"id"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	MenuID       string    `json:"menu_id"`
	MenuName     string    `json:"menu_name"`
	Name         string    `json:"name"`
	Link         string    `json:"link"`
	Image        string    `json:"image"`
	Icon         string    `json:"icon"`
	Active       *bool     `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
}

type MenuDetailRequest struct {
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	MenuID     string `json:"menu_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name       string `json:"name" validate:"required,lowercase,max=200" example:"master produk"`
	Link       string `json:"link"  validate:"required" example:"/master-produk"`
	Image      string `json:"image" example:""`
	Icon       string `json:"icon" example:""`
	Active     *bool  `json:"active" validate:"boolean" example:"true"`
}
