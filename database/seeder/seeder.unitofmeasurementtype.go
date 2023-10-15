package seeder

import (
	"log"
	"time"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeederUnitOfMeasurementType(db *gorm.DB) {
	configs.IsSeederRunning = constants.TRUE_VALUE

	uomTypes := []models.UnitOfMeasurementType{
		{
			ID:         "32838ab3-6773-4db1-b17d-b562eec8a117",
			MerchantID: "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:       "weight",
			Active:     &constants.TRUE_VALUE,
			CreatedAt:  time.Now().Local(),
		},
		{
			ID:         "d4f4d518-ff60-4e59-871d-d49342756416",
			MerchantID: "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:       "volume",
			Active:     &constants.TRUE_VALUE,
			CreatedAt:  time.Now().Local(),
		},
		{
			ID:         "8f15c62e-c8b0-419e-8855-fc8daeda8486",
			MerchantID: "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:       "count",
			Active:     &constants.TRUE_VALUE,
			CreatedAt:  time.Now().Local(),
		},
	}

	for _, uomTypeData := range uomTypes {
		var uomType models.UnitOfMeasurementType // Buat objek uomType baru untuk setiap iterasi

		checkUOM := db.Debug().Where("merchant_id = ? AND name = ?", uomTypeData.MerchantID, uomTypeData.Name).First(&uomType)

		if checkUOM.RowsAffected == 0 {
			if err := db.Create(&uomTypeData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	configs.IsSeederRunning = constants.FALSE_VALUE

	log.Println("Seeder master Unit Of Measurement Type successfully")
}
