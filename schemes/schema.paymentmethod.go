package schemes

import "time"

type PaymentMethod struct {
	ID                string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID        string `json:"merchant_id" validate:"required,uuid"`
	PaymentCategoryID string `json:"payment_category_id" validate:"uuid,required" example:"08d72c5e-2aa1-4c86-9525-0b35199cbc06"`
	// Input with Lowercase
	Name          string `json:"name" validate:"required,lowercase,max=200" example:"tunai"`
	AccountNumber string `json:"account_number" validate:"numeric" example:"787567465"`
	Logo          string `json:"logo"`
	Active        *bool  `json:"active" validate:"boolean" example:"true"`
	Page          int    `json:"page"`
	PerPage       int    `json:"perpage"`
	Sort          string `json:"sort"`
}

type GetPaymentMethod struct {
	ID                  string    `json:"id"`
	MerchantID          string    `json:"merchant_id"`
	MerchantName        string    `json:"merchant_name"`
	PaymentCategoryID   string    `json:"payment_category_id"`
	PaymentCategoryName string    `json:"payment_category_name"`
	Name                string    `json:"name"`
	AccountNumber       string    `json:"account_number"`
	Logo                string    `json:"logo"`
	Active              *bool     `json:"active"`
	CreatedAt           time.Time `json:"created_at"`
}

type PaymentMethodRequest struct {
	MerchantID        string `json:"merchant_id" validate:"required,uuid" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	PaymentCategoryID string `json:"payment_category_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	// Input with Lowercase
	Name          string `json:"name" validate:"required,lowercase,max=200" example:"tunai"`
	AccountNumber string `json:"account_number" validate:"numeric" example:"787567465"`
	Logo          string `json:"logo"`
	Active        *bool  `json:"active" validate:"boolean" example:"true"`
}
