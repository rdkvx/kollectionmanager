package controllers

import (
	"kollectionmanager/m/models"
	"strings"

	"gorm.io/gorm"
)

func CreateGame(game models.Game, db *gorm.DB) error {
	game.Name = strings.ToLower(game.Name)
	game.Deleted = false
	result := db.Create(&game)

	if result.Error != nil {
        return result.Error
    }

	return nil
}