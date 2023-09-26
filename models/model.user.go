package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/pkg"
)

type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Name      string    `json:"name" gorm:"type:varchar;  not null"`
	Email     string    `json:"email" gorm:"type:varchar; unique; not null"`
	Password  string    `json:"password" gorm:"type:varchar; not null"`
	Role      Role      `json:"role" gorm:"foreignkey:RoleID"`
	RoleID    string    `json:"role_id" gorm:"type:varchar; not null"`
	Active    *bool     `json:"active" gorm:"type:boolean; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (User) TableName() string {
	return "master.users"
}

func (m *User) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Password = pkg.HashPassword(m.Password)
	m.Active = &constants.TRUE_VALUE
	m.CreatedAt = time.Now()
	return nil
}

func (m *User) BeforeUpdate(db *gorm.DB) error {
	m.Password = pkg.HashPassword(m.Password)
	m.UpdatedAt = time.Now()
	return nil
}
