package seeder

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeederUnitOfMeasurement(db *gorm.DB) {
	uoms := []models.UnitOfMeasurement{
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "kilogram",
			Symbol:           "kg",
			UOMTypeID:        "32838ab3-6773-4db1-b17d-b562eec8a117",
			ConversionFactor: 1.0,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "gram",
			Symbol:           "g",
			UOMTypeID:        "32838ab3-6773-4db1-b17d-b562eec8a117",
			ConversionFactor: 0.001,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "milligram",
			Symbol:           "mg",
			UOMTypeID:        "32838ab3-6773-4db1-b17d-b562eec8a117",
			ConversionFactor: 0.000001,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "pound",
			Symbol:           "lb",
			UOMTypeID:        "32838ab3-6773-4db1-b17d-b562eec8a117",
			ConversionFactor: 0.453592,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "liter",
			Symbol:           "L",
			UOMTypeID:        "d4f4d518-ff60-4e59-871d-d49342756416",
			ConversionFactor: 1.0,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "milliliter",
			Symbol:           "mL",
			UOMTypeID:        "d4f4d518-ff60-4e59-871d-d49342756416",
			ConversionFactor: 0.001,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "cubic meter",
			Symbol:           "m3",
			UOMTypeID:        "d4f4d518-ff60-4e59-871d-d49342756416",
			ConversionFactor: 1000.0,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "box",
			Symbol:           "box",
			UOMTypeID:        "8f15c62e-c8b0-419e-8855-fc8daeda8486",
			ConversionFactor: 1.0,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
		{
			ID:               uuid.NewString(),
			MerchantID:       "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:             "piece",
			Symbol:           "pcs",
			UOMTypeID:        "8f15c62e-c8b0-419e-8855-fc8daeda8486",
			ConversionFactor: 1.0,
			Active:           &constants.TRUE_VALUE,
			CreatedAt:        time.Now().Local(),
		},
	}

	for _, uomData := range uoms {
		var uom models.UnitOfMeasurement // Buat objek uom baru untuk setiap iterasi

		checkUOM := db.Debug().Where("merchant_id = ? AND uom_type_id = ? AND name = ?", uomData.MerchantID, uomData.UOMTypeID, uomData.Name).First(&uom)

		if checkUOM.RowsAffected == 0 {
			// Hapus ID dari objek uom jika ada
			uomData.ID = ""

			if err := db.Create(&uomData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("Seeder master Unit Of Measurement successfully")
}
