package migration

import (
	"lobmto-echo-example/infrastructure/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Word{}); err != nil {
		return err
	}

	return nil
}
