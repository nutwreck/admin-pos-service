package models

import (
	"time"

	"github.com/nutwreck/admin-pos-service/constants"
	"gorm.io/gorm"
)

type Supplier struct {
	ID          uint64    `json:"id" gorm:"primary_key;autoIncrement"`
	Name        string    `json:"name" gorm:"type:varchar; not null"`
	Phone       string    `json:"phone" gorm:"type:varchar; unique; not null"`
	Address     string    `json:"address" gorm:"type:text; not null"`
	Merchant    Merchant  `json:"merchant" gorm:"foreignkey:MerchantID"`
	MerchantID  string    `json:"merchant_id" gorm:"type:varchar; not null"`
	Outlet      Outlet    `json:"outlet" gorm:"Foreignkey:OutletID;"`
	OutletID    string    `json:"outlet_id" gorm:"index; not null"`
	Description string    `json:"description" gorm:"type:text;"`
	Active      *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Supplier) TableName() string {
	return "master.suppliers"
}

func (m *Supplier) BeforeCreate(db *gorm.DB) error {
	m.Active = &constants.TRUE_VALUE
	m.CreatedAt = time.Now()
	return nil
}

func (m *Supplier) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
