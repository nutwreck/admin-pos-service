package seeder

import (
	"log"
	"time"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeedMerchant(db *gorm.DB) {
	configs.IsSeederRunning = constants.TRUE_VALUE

	merchants := []models.Merchant{
		{
			ID:          "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:        "merchant-master",
			Phone:       "085826046069",
			Address:     "Panjang, Ambarawa, KAB. Semarang",
			Logo:        "DEV_20231005213843_b11a6a0a6bc083aa498593897a6d8585.jpg",
			Description: "Ini adalah toko untuk acuan semua toko yang akan terdaftar",
			Active:      &constants.TRUE_VALUE,
			CreatedAt:   time.Now().Local(),
		},
	}

	for _, merchantData := range merchants {
		var merchant models.Merchant

		checkMerchant := db.Debug().Where("name = ?", merchantData.Name).First(&merchant)

		if checkMerchant.RowsAffected == 0 {
			if err := db.Create(&merchantData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	configs.IsSeederRunning = constants.FALSE_VALUE

	log.Println("Seeder merchant successfully")
}
