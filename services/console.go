package services

import (
	"fmt"
	"kollectionmanager/m/models"
	"kollectionmanager/m/models/dto"

	"gorm.io/gorm"
)

func CreateConsole(console dto.Console, db *gorm.DB) error {

	result := db.Create(&console)

	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Console created")

	return nil
}

func GetConsoles(db *gorm.DB) ([]models.Console, error) {
	var console []models.Console

    result := db.Find(&console)

    if result.Error != nil {
        return nil, result.Error
    }

    return console, nil
}
