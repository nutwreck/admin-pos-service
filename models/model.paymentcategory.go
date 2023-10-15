package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
)

type PaymentCategory struct {
	ID         string    `json:"id" gorm:"primary_key"`
	Merchant   Merchant  `json:"merchant" gorm:"Foreignkey:MerchantID;"`
	MerchantID string    `json:"merchant_id" gorm:"index; not null"`
	Name       string    `json:"name" gorm:"type:varchar; not null"`
	Active     *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (PaymentCategory) TableName() string {
	return "master.payment_categorys"
}

func (m *PaymentCategory) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *PaymentCategory) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
