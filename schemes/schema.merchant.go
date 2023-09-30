package schemes

import "time"

type Merchant struct {
	ID          string    `json:"id" validate:"required,uuid" format:"uuid"`
	Name        string    `json:"name" validate:"required,lowercase"`
	Phone       string    `json:"phone" validate:"required,numeric"`
	Address     string    `json:"address" validate:"required,max=1000"`
	Logo        string    `json:"logo"`
	Description string    `json:"description" validate:"max=1000"`
	CreatedAt   time.Time `json:"created_at"`
	Active      *bool     `json:"active" validate:"boolean" example:"true"`
	Page        int       `json:"page"`
	PerPage     int       `json:"perpage"`
	Sort        string    `json:"sort"`
}

type GetMerchant struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Address     string    `json:"address"`
	Logo        string    `json:"logo"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Active      *bool     `json:"active"`
}
