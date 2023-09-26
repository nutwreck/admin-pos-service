package seeder

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nutwreck/admin-pos-service/constants"
	"github.com/nutwreck/admin-pos-service/models"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) {
	roles := []models.Role{
		{ID: uuid.NewString(), Name: "root", Type: "sys", Active: &constants.TRUE_VALUE, CreatedAt: time.Now()},
		{ID: uuid.NewString(), Name: "superadmin", Type: "user", Active: &constants.TRUE_VALUE, CreatedAt: time.Now()},
	}

	for _, roleData := range roles {
		var role models.Role // Buat objek role baru untuk setiap iterasi

		checkRole := db.Debug().Where("name = ? AND type = ?", roleData.Name, roleData.Type).First(&role)

		if checkRole.RowsAffected == 0 {
			// Hapus ID dari objek role jika ada
			roleData.ID = ""

			if err := db.Create(&roleData).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	log.Println("Seeder master role successfully")
}
