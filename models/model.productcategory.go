package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/constants"
)

type ProductCategory struct {
	ID         uint64    `json:"id" gorm:"primary_key;autoIncrement"`
	Merchant   Merchant  `json:"merchant" gorm:"foreignkey:MerchantID"`
	MerchantID string    `json:"merchant_id" gorm:"type:varchar; not null"`
	Outlet     Outlet    `json:"outlet" gorm:"foreignkey:OutletID"`
	OutletID   string    `json:"outlet_id" gorm:"type:varchar; not null"`
	Name       string    `json:"name" gorm:"type:varchar; not null"`
	Active     *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (ProductCategory) TableName() string {
	return "master.product_categorys"
}

func (m *ProductCategory) BeforeCreate(db *gorm.DB) error {
	m.Active = &constants.TRUE_VALUE
	m.CreatedAt = time.Now()
	return nil
}

func (m *ProductCategory) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
