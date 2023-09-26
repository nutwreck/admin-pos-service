package schemas

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func CreateMasterSchema(db *gorm.DB) error {
	if err := db.Exec("CREATE SCHEMA IF NOT EXISTS master").Error; err != nil {
		logrus.Info("Schema master creation failed")
		return err
	}
	return nil
}
