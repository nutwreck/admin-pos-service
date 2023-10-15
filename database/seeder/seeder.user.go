package seeder

import (
	"log"
	"time"

	"github.com/nutwreck/admin-pos-service/configs"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeedUser(db *gorm.DB) {
	configs.IsSeederRunning = constants.TRUE_VALUE

	users := []models.User{
		{
			ID:         "b4305629-ae51-4837-ab90-02c6498b3bff",
			Name:       "superadmin",
			Email:      "pos.superadmin@digy.com",
			Password:   "$2a$10$tCL0HrW8IrHtkoflpsksfeVET/loibWhJyPizAE6VL44GeyoyxU7q",
			MerchantID: "81c0b615-d575-4d30-a81a-6b8db70fd4e0",
			RoleID:     "01f33858-0cf9-45eb-9e1f-c6a26ca759c4",
			Active:     &constants.TRUE_VALUE,
			CreatedAt:  time.Now().Local(),
		},
	}

	for _, userData := range users {
		var user models.User

		checkUser := db.Debug().Where("email = ?", userData.Email).First(&user)

		if checkUser.RowsAffected == 0 {
			if err := db.Create(&userData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	configs.IsSeederRunning = constants.FALSE_VALUE

	log.Println("Seeder user successfully")
}
