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

func GetConsoles(c *fiber.Ctx, db *gorm.DB) ([]models.Console, error) {
	var console []models.Console

	result := db.Find(&console)

	if result.Error != nil {
		return nil, result.Error
	}

	return console, nil
}

func GetConsoleByName(name string, db *gorm.DB) (dto.ConsoleGet, error) {
	name = strings.ToLower(name)

	var consoleRaw models.Console
	var console dto.ConsoleGet

	result := db.Where(utils.FilterByName, name).Where(utils.FilterByDeleted, false).First(&consoleRaw)
	if result.Error != nil {
		return console, result.Error
	}

	if consoleRaw.Deleted {
		return console, errors.New(utils.FailedTo("find", "console", name))
	}

	console = utils.ConsoleModelToDTO(consoleRaw)

	return console, nil
}

func CreateConsole(consoleRaw dto.ConsolePost, db *gorm.DB) error {

	console, err := utils.ConsoleDTOToModel(consoleRaw)
	if err!= nil {
        return err
    }

	result := db.Create(&console)
	if result.Error != nil {
		return result.Error
	}

	return nil
}


func UpdateConsole(name string, consoleRaw dto.ConsolePost, db *gorm.DB) error {
	name = strings.ToLower(name)

    var console models.Console

    result := db.Where(utils.FilterByName, name).Where(utils.FilterByDeleted, false).First(&console)
    if result.Error != nil {
        return result.Error
    }

    validatedConsole, err := utils.ValidateConsoleData(consoleRaw, console)
	if err != nil {
		return err
	}

    result = db.Save(&validatedConsole)
    if result.Error != nil {
        return result.Error
    }

	return nil
}

func DeleteConsole(name string, db *gorm.DB) error {
	name = strings.ToLower(name)

    var console models.Console

    result := db.Where(utils.FilterByName, name).Where(utils.FilterByDeleted, false).First(&console)
    if result.Error != nil {
        return result.Error
    }

    console.Deleted = true

    result = db.Save(&console)
    if result.Error != nil {
        return result.Error
    }

    return nil
}