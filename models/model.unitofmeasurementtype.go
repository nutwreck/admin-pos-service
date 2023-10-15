package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
)

type UnitOfMeasurementType struct {
	ID         string    `json:"id" gorm:"primary_key"`
	Merchant   Merchant  `json:"merchant" gorm:"foreignkey:MerchantID"`
	MerchantID string    `json:"merchant_id" gorm:"type:varchar; not null"`
	Name       string    `json:"name" gorm:"type:varchar; not null"`
	Active     *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (UnitOfMeasurementType) TableName() string {
	return "master.unit_of_measurement_types"
}

func (m *UnitOfMeasurementType) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *UnitOfMeasurementType) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
