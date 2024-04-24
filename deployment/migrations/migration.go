package migrations

import (
	"kollectionmanager/m/models"

	"gorm.io/gorm"
)

func MigrateIfExists(db *gorm.DB) {
	
	db.AutoMigrate(&models.Manufacturer{})
	db.AutoMigrate(&models.Console{})
	db.AutoMigrate(&models.Developer{})
	db.AutoMigrate(&models.Game{})
}
