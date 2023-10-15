package seeder

import (
	"log"
	"time"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeedOutlet(db *gorm.DB) {
	configs.IsSeederRunning = constants.TRUE_VALUE

	outlets := []models.Outlet{
		{
			ID:          "4e769a02-0214-4277-90d0-bdf7f7b7c064",
			Name:        "outlet-master",
			Phone:       "085826046069",
			Address:     "Panjang, Ambarawa, KAB. Semarang",
			MerchantID:  "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Description: "Ini adalah outlet untuk acuan semua outlet yang akan terdaftar",
			IsPrimary:   &constants.TRUE_VALUE,
			Active:      &constants.TRUE_VALUE,
			CreatedAt:   time.Now().Local(),
		},
	}

	for _, outletData := range outlets {
		var outlet models.Outlet

		checkOutlet := db.Debug().Where("name = ?", outletData.Name).First(&outlet)

		if checkOutlet.RowsAffected == 0 {
			if err := db.Create(&outletData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	configs.IsSeederRunning = constants.FALSE_VALUE

	log.Println("Seeder outlet successfully")
}
