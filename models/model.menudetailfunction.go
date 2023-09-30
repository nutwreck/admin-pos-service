package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/constants"
)

type MenuDetailFunction struct {
	ID           string     `json:"id" gorm:"primary_key"`
	Menu         Menu       `json:"menu" gorm:"foreignkey:MenuID"`
	MenuID       string     `json:"menu_id" gorm:"type:varchar;  not null"`
	MenuDetail   MenuDetail `json:"menu_detail" gorm:"foreignkey:MenuDetailID"`
	MenuDetailID string     `json:"menu_detail_id" gorm:"type:varchar;  not null"`
	Name         string     `json:"name" gorm:"type:varchar;  not null"`
	Link         string     `json:"link" gorm:"type:varchar; not null"`
	Active       *bool      `json:"active" gorm:"type:boolean; not null"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (MenuDetailFunction) TableName() string {
	return "master.menu_detail_functions"
}

func (m *MenuDetailFunction) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Active = &constants.TRUE_VALUE
	m.CreatedAt = time.Now()
	return nil
}

func (m *MenuDetailFunction) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now()
	return nil
}
