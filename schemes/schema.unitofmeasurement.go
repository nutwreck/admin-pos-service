package schemes

import "time"

type UnitOfMeasurement struct {
	ID         string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	UOMTypeID  string `json:"uom_type_id" validate:"uuid,required" example:"32838ab3-6773-4db1-b17d-b562eec8a117"`
	Symbol     string `json:"symbol" validate:"required" example:"kg"`
	// Input with Lowercase
	Name             string  `json:"name" validate:"required,lowercase,max=200" example:"Kilogram"`
	ConversionFactor float64 `json:"conversion_factor" validate:"required,numeric" example:"0.1"`
	Active           *bool   `json:"active" validate:"boolean" example:"true"`
	Page             int     `json:"page"`
	PerPage          int     `json:"perpage"`
	Sort             string  `json:"sort"`
}

type GetUnitOfMeasurement struct {
	ID               string    `json:"id"`
	MerchantID       string    `json:"merchant_id"`
	MerchantName     string    `json:"merchant_name"`
	UOMTypeID        string    `json:"uom_type_id"`
	UOMNameID        string    `json:"uom_type_name"`
	Name             string    `json:"name"`
	Symbol           string    `json:"symbol"`
	ConversionFactor float64   `json:"conversion_factor"`
	Active           *bool     `json:"active"`
	CreatedAt        time.Time `json:"created_at"`
}

type UnitOfMeasurementRequest struct {
	MerchantID string `json:"merchant_id" validate:"uuid,required" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	UOMTypeID  string `json:"uom_type_id" validate:"uuid,required" example:"32838ab3-6773-4db1-b17d-b562eec8a117"`
	Symbol     string `json:"symbol" validate:"required" example:"kg"`
	// Input with Lowercase
	Name             string  `json:"name" validate:"required,lowercase,max=200" example:"Kilogram"`
	ConversionFactor float64 `json:"conversion_factor" validate:"required,numeric" example:"0.1"`
	Active           *bool   `json:"active" validate:"boolean" example:"true"`
}
