package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"gorm.io/gorm"
)

type UnitOfMeasurement struct {
	ID               string                `json:"id" gorm:"primary_key"`
	Merchant         Merchant              `json:"merchant" gorm:"foreignkey:MerchantID"`
	MerchantID       string                `json:"merchant_id" gorm:"type:varchar; not null"`
	Name             string                `json:"name" gorm:"type:varchar; not null"`
	Symbol           string                `json:"symbol" gorm:"type:varchar; not null"`
	UOMType          UnitOfMeasurementType `json:"uom_type" gorm:"foreignkey:UOMTypeID"`
	UOMTypeID        string                `json:"uom_type_id" gorm:"type:varchar; not null"`
	ConversionFactor float64               `json:"conversion_factor" gorm:"type:double precision; not null; default=0"`
	Active           *bool                 `json:"active" gorm:"type:boolean; not null"`
	CreatedAt        time.Time             `json:"created_at"`
	UpdatedAt        time.Time             `json:"updated_at"`
}

func (UnitOfMeasurement) TableName() string {
	return "master.unit_of_measurements"
}

func (m *UnitOfMeasurement) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *UnitOfMeasurement) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}

//amountInMilligram := amountInGram * gram.ConversionFactor / milligram.ConversionFactor
