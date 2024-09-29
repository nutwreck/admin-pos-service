package schemes

import "time"

type MappingRoleMenuUser struct {
	ID                string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID        string `json:"merchant_id" validate:"required,uuid"`
	UserID            string `json:"user_id" validate:"required,uuid"`
	MappingRoleMenuID string `json:"mapping_role_menu_id" validate:"required,uuid"`
	Active            *bool  `json:"active" validate:"boolean" example:"true"`
	Page              int    `json:"page"`
	PerPage           int    `json:"perpage"`
	Sort              string `json:"sort"`
}

type GetMappingRoleMenuUser struct {
	ID                  string    `json:"id"`
	MerchantID          string    `json:"merchant_id"`
	MerchantName        string    `json:"merchant_name"`
	UserID              string    `json:"user_id"`
	UserName            string    `json:"user_name"`
	MappingRoleMenuID   string    `json:"mapping_role_menu_id"`
	MappingRoleMenuName string    `json:"mapping_role_menu_name"`
	Active              *bool     `json:"active"`
	CreatedAt           time.Time `json:"created_at"`
}

type MappingRoleMenuUserRequest struct {
	MerchantID        string `json:"merchant_id" validate:"required,uuid" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	UserID            string `json:"user_id" validate:"required,uuid" example:"4e769a02-0214-4277-90d0-bdf7f7b7c064"`
	MappingRoleMenuID string `json:"mapping_role_menu_id" validate:"required,uuid" example:"4e769a02-0214-4277-90d0-bdf7f7b7c064"`
	Active            *bool  `json:"active" validate:"boolean" example:"true"`
}
