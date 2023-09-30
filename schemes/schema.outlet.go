package schemes

import "time"

type Outlet struct {
	ID          string    `json:"id" validate:"required,uuid" format:"uuid"`
	Name        string    `json:"name" validate:"required,lowercase"`
	Phone       string    `json:"phone" validate:"required,gte=12"`
	Address     string    `json:"address" validate:"required,max=1000"`
	MerchantID  string    `json:"merchant_id" validate:"required,uuid"`
	Description string    `json:"description" validate:"max=1000"`
	CreatedAt   time.Time `json:"created_at"`
	Active      *bool     `json:"active" validate:"boolean" example:"true"`
	Page        int       `json:"page"`
	PerPage     int       `json:"perpage"`
	Sort        string    `json:"sort"`
}

type GetOutlet struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	MerchantID   string    `json:"merchant_id"`
	MerchantName string    `json:"merchant_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	Active       *bool     `json:"active"`
}

type OutletRequest struct {
	// Input with Lowercase
	Name        string `json:"name" validate:"required,lowercase" example:"cabang 1 warung berkah"`
	Phone       string `json:"phone" validate:"required,gte=12" example:"085768576857"`
	Address     string `json:"address" validate:"required,max=1000" example:"jl. merdeka barat, ngaliyan, kota semarang"`
	MerchantID  string `json:"merchant_id" validate:"required,uuid" example:"550e8400-e29b-41d4-a716-446655440000"`
	Description string `json:"description" validate:"max=1000" example:"isi dengan catatan tentang outlet ini"`
	Active      *bool  `json:"active" validate:"boolean" example:"true"`
}
