package controllers

import (
	"kollectionmanager/m/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetConsoles(c *fiber.Ctx, db *gorm.DB) ([]models.Console, error) {
	var console []models.Console

	result := db.Find(&console)

	if result.Error != nil {
		return nil, result.Error
	}

	return console, nil
}

func CreateConsole(console models.Console, db *gorm.DB) error {
	console.Name = strings.ToLower(console.Name)
	console.Deleted = false

	result := db.Create(&console)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
