package schemes

import "time"

type Product struct {
	ID                   string `json:"id" validate:"uuid" example:"550e8400-e29b-41d4-a716-446655440000" format:"uuid"`
	MerchantID           string `json:"merchant_id" validate:"required,uuid" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	OutletID             string `json:"outlet_id" validate:"required,uuid" example:"4e769a02-0214-4277-90d0-bdf7f7b7c064"`
	ProductCategoryID    string `json:"product_category_id" validate:"required,uuid" example:"8938d34c-f853-44c3-86f4-ca499c82c4f3"`
	ProductCategorySubID string `json:"product_category_sub_id" validate:"required,uuid" example:"8938d34c-f853-44c3-86f4-ca499c82c4f3"`
	Code                 string `json:"code" validate:"required" example:"AQU12"`
	// Input with Lowercase
	Name                string  `json:"name" validate:"required,lowercase,max=200" example:"aqua"`
	Barcode             string  `json:"barcode"`
	CapitalPrice        float64 `json:"capital_price" validate:"required,numeric" example:"4500"`
	SellingPrice        float64 `json:"selling_price" validate:"required,numeric" example:"5500"`
	SupplierID          string  `json:"supplier_id" validate:"uuid" example:"1320754f-613d-43d5-beb8-3e275c5e02ed"`
	UnitOfMeasurementID string  `json:"unit_of_measurement_id" validate:"required,uuid" example:"b8531dc1-d71e-499a-a6c3-ad10ada7c8d1"`
	Image               string  `json:"image"`
	Active              *bool   `json:"active" validate:"boolean" example:"true"`
	Page                int     `json:"page"`
	PerPage             int     `json:"perpage"`
	Sort                string  `json:"sort"`
}

type GetProduct struct {
	ID                      string `json:"id"`
	MerchantID              string `json:"merchant_id"`
	MerchantName            string `json:"merchant_name"`
	OutletID                string `json:"outlet_id"`
	OutletName              string `json:"outlet_name"`
	ProductCategoryID       string `json:"product_category_id"`
	ProductCategoryName     string `json:"product_category_name"`
	ProductCategorySubID    string `json:"product_category_sub_id"`
	ProductCategoryaSubName string `json:"product_category_sub_name"`
	Code                    string `json:"code" validate:"required"`
	// Input with Lowercase
	Name                  string    `json:"name"`
	Barcode               string    `json:"barcode"`
	CapitalPrice          float64   `json:"capital_price"`
	SellingPrice          float64   `json:"selling_price"`
	SupplierID            string    `json:"supplier_id"`
	SupplierName          string    `json:"supplier_name"`
	UnitOfMeasurementID   string    `json:"unit_of_measurement_id"`
	UnitOfMeasurementName string    `json:"unit_of_measurement_name"`
	Image                 string    `json:"image"`
	Active                *bool     `json:"active" validate:"boolean"`
	CreatedAt             time.Time `json:"created_at"`
}

type ProductRequest struct {
	MerchantID           string `json:"merchant_id" validate:"required,uuid" example:"81c0b615-d575-4d30-a81a-6b8db70fd4e0"`
	OutletID             string `json:"outlet_id" validate:"required,uuid" example:"4e769a02-0214-4277-90d0-bdf7f7b7c064"`
	ProductCategoryID    string `json:"product_category_id" validate:"required,uuid" example:"8938d34c-f853-44c3-86f4-ca499c82c4f3"`
	ProductCategorySubID string `json:"product_category_sub_id" validate:"required,uuid" example:"8938d34c-f853-44c3-86f4-ca499c82c4f3"`
	Code                 string `json:"code" validate:"required" example:"AQU12"`
	// Input with Lowercase
	Name                string  `json:"name" validate:"required,lowercase,max=200" example:"aqua"`
	Barcode             string  `json:"barcode" example:""`
	CapitalPrice        float64 `json:"capital_price" validate:"required,numeric" example:"4500"`
	SellingPrice        float64 `json:"selling_price" validate:"required,numeric" example:"5500"`
	SupplierID          string  `json:"supplier_id" example:"1320754f-613d-43d5-beb8-3e275c5e02ed"`
	UnitOfMeasurementID string  `json:"unit_of_measurement_id" validate:"required,uuid" example:"b8531dc1-d71e-499a-a6c3-ad10ada7c8d1"`
	Image               string  `json:"image" example:""`
	Active              *bool   `json:"active" validate:"boolean" example:"true"`
}
