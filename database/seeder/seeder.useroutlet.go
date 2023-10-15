package seeder

import (
	"log"
	"time"

	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeedUserOutlet(db *gorm.DB) {
	userOutlets := []models.UserOutlet{
		{
			UserID:    "b4305629-ae51-4837-ab90-02c6498b3bff",
			OutletID:  "4e769a02-0214-4277-90d0-bdf7f7b7c064",
			Active:    &constants.TRUE_VALUE,
			CreatedAt: time.Now().Local(),
		},
	}

	for _, userOutletData := range userOutlets {
		var userOutlet models.UserOutlet

		checkData := db.Debug().Where("user_id = ? AND outlet_id = ?", userOutletData.UserID, userOutletData.OutletID).First(&userOutlet)

		if checkData.RowsAffected == 0 {
			if err := db.Create(&userOutletData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("Seeder user outlet successfully")
}
