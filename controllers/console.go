package controllers

import (
	"kollectionmanager/m/models"
	"kollectionmanager/m/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GetConsoles(c *fiber.Ctx, db *gorm.DB) ([]models.Console, error) {
	/* consoles := models.Console{}

	if err := c.BodyParser(&consoles); err != nil {
		return err
	} */

	consoles, err := services.GetConsoles(db)
	if err != nil {
		consolenil := []models.Console{}
		return consolenil, err
	}

	return consoles, nil
}
