package seeder

import (
	"log"
	"time"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	configs.IsSeederRunning = constants.TRUE_VALUE

	roles := []models.Role{
		{
			ID:         "01f33858-0cf9-45eb-9e1f-c6a26ca759c4",
			MerchantID: "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:       "superadmin",
			Type:       "sys",
			Active:     &constants.TRUE_VALUE,
			CreatedAt:  time.Now().Local(),
		},
		// {ID: uuid.NewString(), Name: "superadmin", Type: "user", Active: &constants.TRUE_VALUE, CreatedAt: time.Now().Local()},
	}

	for _, roleData := range roles {
		var role models.Role // Buat objek role baru untuk setiap iterasi

		checkRole := db.Debug().Where("name = ? AND type = ? AND merchant_id = ?", roleData.Name, roleData.Type, roleData.MerchantID).First(&role)

		if checkRole.RowsAffected == 0 {
			if err := db.Create(&roleData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	configs.IsSeederRunning = constants.FALSE_VALUE

	log.Println("Seeder master role successfully")
}
