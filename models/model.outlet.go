package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"gorm.io/gorm"
)

type Outlet struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar; not null"`
	Phone       string    `json:"phone" gorm:"type:varchar; unique; not null"`
	Address     string    `json:"address" gorm:"type:text; not null"`
	Merchant    Merchant  `json:"merchant" gorm:"Foreignkey:MerchantID;"`
	MerchantID  string    `json:"merchant_id" gorm:"index; not null"`
	Description string    `json:"description" gorm:"type:text;"`
	IsPrimary   *bool     `json:"is_primary" gorm:"type:boolean; not null"`
	Active      *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Outlet) TableName() string {
	return "master.outlets"
}

func (m *Outlet) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *Outlet) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
