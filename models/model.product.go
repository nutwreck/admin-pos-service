package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
)

type Product struct {
	ID                   string             `json:"id" gorm:"primary_key"`
	Merchant             Merchant           `json:"merchant" gorm:"foreignkey:MerchantID"`
	MerchantID           string             `json:"merchant_id" gorm:"type:varchar; not null"`
	Outlet               Outlet             `json:"outlet" gorm:"foreignkey:OutletID"`
	OutletID             string             `json:"outlet_id" gorm:"type:varchar; not null"`
	ProductCategory      ProductCategory    `json:"product_category" gorm:"foreignkey:ProductCategoryID"`
	ProductCategoryID    string             `json:"product_category_id" gorm:"type:varchar; not null"`
	ProductCategorySub   ProductCategorySub `json:"product_category_sub" gorm:"foreignkey:ProductCategorySubID"`
	ProductCategorySubID string             `json:"product_category_sub_id" gorm:"type:varchar; not null"`
	Code                 string             `json:"code" gorm:"type:varchar; not null"`
	Name                 string             `json:"name" gorm:"type:varchar; not null"`
	Barcode              string             `json:"barcode" gorm:"type:varchar;"`
	CapitalPrice         float64            `json:"capital_price" gorm:"type:double precision; not null; default=0"`
	SellingPrice         float64            `json:"selling_price" gorm:"type:double precision; not null; default=0"`
	SupplierID           string             `json:"supplier_id" gorm:"type:varchar;"`
	UnitOfMeasurement    UnitOfMeasurement  `json:"unit_of_measurement" gorm:"foreignkey:UnitOfMeasurementID"`
	UnitOfMeasurementID  string             `json:"unit_of_measurement_id" gorm:"type:varchar; not null"`
	Image                string             `json:"image" gorm:"type:varchar;"`
	Active               *bool              `json:"active" gorm:"type:boolean; not null"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
}

func (Product) TableName() string {
	return "master.products"
}

func (m *Product) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *Product) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
