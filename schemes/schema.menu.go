package schemes

import (
	"time"
)

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

type GetMenuRelationRaw struct {
	MenuID                   string `json:"menu_id"`
	MenuName                 string `json:"menu_name"`
	MenuActive               *bool  `json:"menu_active"`
	MenuDetailID             string `json:"menu_detail_id"`
	MenuDetailName           string `json:"menu_detail_name"`
	MenuDetailLink           string `json:"menu_detail_link"`
	MenuDetailImage          string `json:"menu_detail_image"`
	MenuDetailIcon           string `json:"menu_detail_icon"`
	MenuDetailActive         *bool  `json:"menu_detail_active"`
	MenuDetailFunctionID     string `json:"menu_detail_function_id"`
	MenuDetailFunctionName   string `json:"menu_detail_function_name"`
	MenuDetailFunctionLink   string `json:"menu_detail_function_link"`
	MenuDetailFunctionActive *bool  `json:"menu_detail_function_active"`
}

type GroupMenuKey struct {
	MenuID   string
	MenuName string
}

type GetMenuRelation struct {
	ID         string           `json:"id"`
	Name       string           `json:"label_group"`
	Active     *bool            `json:"active"`
	ListDetail []ListMenuDetail `json:"list_detail"`
}

type ListMenuDetail struct {
	ID                 string                   `json:"id"`
	Name               string                   `json:"name"`
	Link               string                   `json:"link"`
	Image              string                   `json:"image"`
	Icon               string                   `json:"icon"`
	Active             *bool                    `json:"active"`
	ListDetailFunction []ListMenuDetailFunction `json:"list_function"`
}

type ListMenuDetailFunction struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Link   string `json:"link"`
	Active *bool  `json:"active"`
}
