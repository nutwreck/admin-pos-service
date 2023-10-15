package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
)

type Role struct {
	ID         string    `json:"id" gorm:"primary_key"`
	Merchant   Merchant  `json:"merchant" gorm:"foreignkey:MerchantID"`
	MerchantID string    `json:"merchant_id" gorm:"type:varchar; not null"`
	Name       string    `json:"name" gorm:"type:varchar;  not null"`
	Type       string    `json:"type" gorm:"type:varchar; not null"`
	Active     *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Role) TableName() string {
	return "master.roles"
}

func (m *Role) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *Role) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
