package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"gorm.io/gorm"
)

type Sales struct {
	ID          string    `json:"id" gorm:"primary_key"`
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

func (Sales) TableName() string {
	return "master.sales"
}

func (m *Sales) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *Sales) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
