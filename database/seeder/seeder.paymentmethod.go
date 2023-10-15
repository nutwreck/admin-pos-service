package seeder

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeederPaymentMethod(db *gorm.DB) {
	paymentMethods := []models.PaymentMethod{
		{
			ID:                uuid.NewString(),
			MerchantID:        "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			Name:              "tunai",
			PaymentCategoryID: "08d72c5e-2aa1-4c86-9525-0b35199cbc06",
			AccountNumber:     constants.EMPTY_VALUE,
			Logo:              constants.EMPTY_VALUE,
			Active:            &constants.TRUE_VALUE,
			CreatedAt:         time.Now().Local(),
		},
	}

	for _, paymentMethodData := range paymentMethods {
		var paymentMethod models.PaymentMethod

		checkData := db.Debug().Where("merchant_id = ? AND payment_category_id = ? AND name = ?", paymentMethodData.MerchantID, paymentMethodData.PaymentCategoryID, paymentMethodData.Name).First(&paymentMethod)

		if checkData.RowsAffected == 0 {
			// Hapus ID dari objek payment method jika ada
			paymentMethodData.ID = ""

			if err := db.Create(&paymentMethodData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("Seeder master Payment Method successfully")
}
