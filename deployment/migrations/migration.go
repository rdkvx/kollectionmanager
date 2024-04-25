package migrations

import (
	"fmt"
	"kollectionmanager/m/models"
	"kollectionmanager/m/utils"

	"gorm.io/gorm"
)

func MigrateIfExists(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.Developer{}) {
		db.AutoMigrate(&models.Manufacturer{})
		db.AutoMigrate(&models.Console{})
		db.AutoMigrate(&models.Developer{})
		db.AutoMigrate(&models.Game{})
		fmt.Println(utils.MigrationsSuccess)
	}
}
