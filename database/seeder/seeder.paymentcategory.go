package seeder

import (
	"log"
	"time"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeederPaymentCategory(db *gorm.DB) {
	configs.IsSeederRunning = constants.TRUE_VALUE

	paymentCategorys := []models.PaymentCategory{
		{
			ID:         "08d72c5e-2aa1-4c86-9525-0b35199cbc06",
			MerchantID: "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:       "tunai",
			Active:     &constants.TRUE_VALUE,
			CreatedAt:  time.Now().Local()},
	}

	for _, paymentCategoryData := range paymentCategorys {
		var paymentCategory models.PaymentCategory

		checkData := db.Debug().Where("merchant_id = ? AND name = ?", paymentCategoryData.MerchantID, paymentCategoryData.Name).First(&paymentCategory)

		if checkData.RowsAffected == 0 {
			if err := db.Create(&paymentCategoryData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	configs.IsSeederRunning = constants.FALSE_VALUE

	log.Println("Seeder master Payment Category successfully")
}
