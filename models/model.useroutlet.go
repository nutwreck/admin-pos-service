package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
)

type UserOutlet struct {
	ID        string    `json:"id" gorm:"primary_key"`
	User      User      `json:"user" gorm:"foreignkey:UserID"`
	UserID    string    `json:"user_id" gorm:"type:varchar;  not null"`
	Outlet    Outlet    `json:"outlet" gorm:"foreignkey:OutletID"`
	OutletID  string    `json:"outlet_id" gorm:"type:varchar; not null"`
	Active    *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserOutlet) TableName() string {
	return "master.user_outlets"
}

func (m *UserOutlet) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *UserOutlet) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
