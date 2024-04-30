package migrations

import (
	"fmt"
	"kollectionmanager/m/models"
	"kollectionmanager/m/utils"
	"os"

	"gorm.io/gorm"
)

func MigrateIfExists(db *gorm.DB) {

	if os.Getenv("MIGRATIONUP") == "TRUE" {
		db.AutoMigrate(&models.Manufacturer{})
		db.AutoMigrate(&models.Console{})
		db.AutoMigrate(&models.Developer{})
		db.AutoMigrate(&models.Game{})
		fmt.Println(utils.MigrationsSuccess)
	}
}
