package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
)

type PaymentMethod struct {
	ID                string          `json:"id" gorm:"primary_key"`
	Merchant          Merchant        `json:"merchant" gorm:"Foreignkey:MerchantID;"`
	MerchantID        string          `json:"merchant_id" gorm:"index; not null"`
	PaymentCategory   PaymentCategory `json:"payment_category" gorm:"foreignkey:PaymentCategoryID"`
	PaymentCategoryID string          `json:"payment_category_id" gorm:"type:varchar; not null"`
	Name              string          `json:"name" gorm:"type:varchar; not null"`
	AccountNumber     string          `json:"account_number" gorm:"type:varchar;"`
	Logo              string          `json:"logo" gorm:"type:varchar;"`
	Active            *bool           `json:"active" gorm:"type:boolean; not null"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
}

func (PaymentMethod) TableName() string {
	return "master.payment_methods"
}

func (m *PaymentMethod) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *PaymentMethod) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
