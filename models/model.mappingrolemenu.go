package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
)

type MappingRoleMenu struct {
	ID                   string             `json:"id" gorm:"primary_key"`
	Name                 string             `json:"name" gorm:"type:varchar; not null"`
	Merchant             Merchant           `json:"merchant" gorm:"foreignkey:MerchantID"`
	MerchantID           string             `json:"merchant_id" gorm:"type:varchar; not null"`
	Role                 Role               `json:"role" gorm:"foreignkey:RoleID"`
	RoleID               string             `json:"role_id" gorm:"type:varchar;  not null"`
	Menu                 Menu               `json:"menu" gorm:"foreignkey:MenuID"`
	MenuID               string             `json:"menu_id" gorm:"type:varchar;  not null"`
	MenuDetail           MenuDetail         `json:"menu_detail" gorm:"foreignkey:MenuDetailID"`
	MenuDetailID         string             `json:"menu_detail_id" gorm:"type:varchar;  not null"`
	MenuDetailFunction   MenuDetailFunction `json:"menu_detail_function" gorm:"foreignkey:MenuDetailFunctionID"`
	MenuDetailFunctionID string             `json:"menu_detail_function_id" gorm:"type:varchar;  not null"`
	Active               *bool              `json:"active" gorm:"type:boolean; not null"`
	CreatedAt            time.Time          `json:"created_at"`
	UpdatedAt            time.Time          `json:"updated_at"`
}

func (MappingRoleMenu) TableName() string {
	return "master.mapping_role_menus"
}

func (m *MappingRoleMenu) BeforeCreate(db *gorm.DB) error {
	if !configs.IsSeederRunning {
		m.ID = uuid.NewString()
		m.Active = &constants.TRUE_VALUE
		m.CreatedAt = time.Now().Local()
	}
	return nil
}

func (m *MappingRoleMenu) BeforeUpdate(db *gorm.DB) error {
	m.UpdatedAt = time.Now().Local()
	return nil
}
