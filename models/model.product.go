package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
)

type Product struct {
	ID                    uint64             `json:"id" gorm:"primary_key;autoIncrement"`
	Merchant              Merchant           `json:"merchant" gorm:"foreignkey:MerchantID"`
	MerchantID            string             `json:"merchant_id" gorm:"type:varchar; not null"`
	Outlet                Outlet             `json:"outlet" gorm:"foreignkey:OutletID"`
	OutletID              string             `json:"outlet_id" gorm:"type:varchar; not null"`
	ProductCategory       ProductCategory    `json:"product_category" gorm:"foreignkey:ProductCategoryID"`
	ProductCategoryID     uint64             `json:"product_category_id" gorm:"type:int; not null"`
	ProductCategorySub    ProductCategorySub `json:"product_category_sub" gorm:"foreignkey:ProductCategoryaSubID"`
	ProductCategoryaSubID uint64             `json:"product_category_sub_id" gorm:"type:int; not null"`
	Code                  string             `json:"code" gorm:"type:varchar; not null"`
	Name                  string             `json:"name" gorm:"type:varchar; not null"`
	Barcode               string             `json:"barcode" gorm:"type:varchar;"`
	CapitalPrice          float64            `json:"capital_price" gorm:"type:double precision; not null; default=0"`
	SellingPrice          float64            `json:"selling_price" gorm:"type:double precision; not null; default=0"`
	Supplier              Supplier           `json:"supplier" gorm:"foreignkey:SupplierID"`
	SupplierID            uint64             `json:"supplier_id" gorm:"type:int;"`
	UnitOfMeasurement     UnitOfMeasurement  `json:"unit_of_measurement" gorm:"foreignkey:UnitOfMeasurementID"`
	UnitOfMeasurementID   string             `json:"unit_of_measurement_id" gorm:"type:varchar; not null"`
	Active                *bool              `json:"active" gorm:"type:boolean; not null"`
	CreatedAt             time.Time          `json:"created_at"`
	UpdatedAt             time.Time          `json:"updated_at"`
}

func (Product) TableName() string {
	return "master.products"
}

func (m *Product) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *Product) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
