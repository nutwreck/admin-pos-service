package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/constants"
	"gorm.io/gorm"
)

type Merchant struct {
	ID          string    `json:"id" gorm:"primary_key"`
	Name        string    `json:"name" gorm:"type:varchar; not null"`
	Phone       string    `json:"phone" gorm:"type:varchar; unique; not null"`
	Address     string    `json:"address" gorm:"type:text; not null"`
	Logo        string    `json:"logo" gorm:"type:varchar; not null"`
	Description string    `json:"description" gorm:"type:text;"`
	Active      *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Merchant) TableName() string {
	return "master.merchants"
}

func (m *Merchant) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.New().String()
	m.Active = &constants.TRUE_VALUE
	m.CreatedAt = time.Now().Local()
	return nil
}

func (m *Merchant) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
