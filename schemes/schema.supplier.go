package schemes

import "time"

type Supplier struct {
	ID          string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	Name        string `json:"name" validate:"required,lowercase"`
	Phone       string `json:"phone" validate:"required,numeric"`
	Address     string `json:"address" validate:"required,max=1000"`
	Description string `json:"description" validate:"max=1000"`
	MerchantID  string `json:"merchant_id" validate:"required,uuid"`
	OutletID    string `json:"outlet_id" validate:"required,uuid"`
	Active      *bool  `json:"active" validate:"boolean" example:"true"`
	Page        int    `json:"page"`
	PerPage     int    `json:"perpage"`
	Sort        string `json:"sort"`
}

type GetSupplier struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	Description  string    `json:"description"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	OutletID     string    `json:"outlet_id"`
	OutletName   string    `json:"outlet_name"`
	Active       *bool     `json:"active"`
	CreatedAt    time.Time `json:"created_at"`
}

type SupplierRequest struct {
	Name        string `json:"name" validate:"required,lowercase" example:"cv makmur sentosa"`
	Phone       string `json:"phone" validate:"required,gte=12,numeric" example:"087987875765"`
	Address     string `json:"address" validate:"required,max=1000" example:"JL. Pahlawan, Ngaliyan, Semarang"`
	Description string `json:"description" validate:"max=1000" example:"Supplier beras"`
	MerchantID  string `json:"merchant_id" validate:"required,uuid" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	OutletID    string `json:"outlet_id" validate:"required,uuid" example:"4e769a02-0214-4277-90d0-bdf7f7b7c064"`
	Active      *bool  `json:"active" validate:"boolean" example:"true"`
}
