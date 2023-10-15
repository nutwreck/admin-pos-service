package schemes

import "time"

type UserOutlet struct {
	ID         string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	UserID     string `json:"user_id" validate:"required,uuid"`
	OutletID   string `json:"outlet_id" validate:"required,uuid"`
	Active     *bool  `json:"active" validate:"boolean" example:"true"`
	Page       int    `json:"page"`
	PerPage    int    `json:"perpage"`
	Sort       string `json:"sort"`
}

type GetUserOutlet struct {
	ID                string    `json:"id"`
	UserID            string    `json:"user_id"`
	UserName          string    `json:"user_name"`
	OutletID          string    `json:"outlet_id"`
	OutletName        string    `json:"outlet_name"`
	OutletPhone       string    `json:"outlet_phone"`
	OutletAddress     string    `json:"outlet_address"`
	OutletDescription string    `json:"outlet_description"`
	OutletActive      *bool     `json:"outlet_active"`
	OutletCreatedAt   time.Time `json:"outlet_created_at"`
}

type GetAllUserOutlet struct {
	ID           string    `json:"id"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	OutletID     string    `json:"outlet_id"`
	OutletName   string    `json:"outlet_name"`
	Active       *bool     `json:"outlet_active"`
	CreatedAt    time.Time `json:"outlet_created_at"`
}

type UserOutletRequest struct {
	UserID   string `json:"user_id" validate:"required,uuid" example:"b4305629-ae51-4837-ab90-02c6498b3bff"`
	OutletID string `json:"outlet_id" validate:"required,uuid" example:"4e769a02-0214-4277-90d0-bdf7f7b7c064"`
	Active   *bool  `json:"active" validate:"boolean" example:"true"`
}
