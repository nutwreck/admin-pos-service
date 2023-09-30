package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/constants"
	"gorm.io/gorm"
)

type Outlet struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar; not null"`
	Phone       uint64    `json:"phone" gorm:"type:bigint; unique; not null"`
	Address     string    `json:"address" gorm:"type:text; not null"`
	Merchant    Merchant  `json:"merchant" gorm:"Foreignkey:MerchantID;"`
	MerchantID  string    `json:"merchant_id" gorm:"index; not null"`
	Description string    `json:"description" gorm:"type:text;"`
	Active      *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Outlet) TableName() string {
	return "master.outlets"
}

func (m *Outlet) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Active = &constants.TRUE_VALUE
	m.CreatedAt = time.Now()
	return nil
}

func (m *Outlet) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
