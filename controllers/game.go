package controllers

import (
	"errors"
	"kollectionmanager/m/models"
	"kollectionmanager/m/models/dto"
	"kollectionmanager/m/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetGames(c *fiber.Ctx, db *gorm.DB) ([]dto.GameGet, error) {
	var gamesRaw []models.Game

    result := db.Find(&gamesRaw)

    if result.Error!= nil {
        return nil, result.Error
    }

	var games []dto.GameGet

	for _, game := range gamesRaw {
		if !game.Deleted {
            games = append(games, utils.GameModelToDTO(game))
        }
	}

    return games, nil
}

func GetGameByName(name string, db *gorm.DB) (dto.GameGet, error){
	name = strings.ToLower(name)

	var gameRaw models.Game
	var game dto.GameGet

	result := db.Where(utils.FilterByName, name).Where(utils.FilterByDeleted, false).First(&gameRaw)
	if result.Error != nil {
		return game, result.Error
	}

	if gameRaw.Deleted {
        return game, errors.New(utils.FailedTo("find", "game", name))
    }

	game = utils.GameModelToDTO(gameRaw)
	return game, nil
}

func CreateGame(gameRaw dto.GamePost, db *gorm.DB) error {
	game, err := utils.GameDTOToModel(gameRaw)
	if err != nil {
		return err
	}

	result := db.Create(&game)

	if result.Error != nil {
        return result.Error
    }

	return nil
}

func UpdateGame(name string, gameRaw dto.GamePost, db *gorm.DB) error{
	name = strings.ToLower(name)
	game := models.Game{}

	result := db.Where(utils.FilterByName, name).Where(utils.FilterByDeleted, false).First(game)
	if result.Error != nil {
		return result.Error
	}

	updatedGame, err := utils.ValidateGameData(gameRaw, game)
	if err != nil {
        return err
    }

	result = db.Save(&updatedGame)
	if result.Error != nil {
        return result.Error
    }

    return nil
}

func SoftDeleteGameByName(name string, db *gorm.DB) error {
	name = strings.ToLower(name)

    var game models.Game

    result := db.Where(utils.FilterByName, name).Where(utils.FilterByDeleted, false).First(&game)
    if result.Error != nil {
        return result.Error
    }

	if game.Deleted {
        return errors.New("record not found")
    }

    game.Deleted = true

    result = db.Save(&game)
    if result.Error != nil {
        return result.Error
    }

    return nil
}