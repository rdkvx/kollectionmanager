package routes

import (
	"kollectionmanager/m/controllers"
	"kollectionmanager/m/models/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ManufacturerRoutes(app *fiber.App, db *gorm.DB) {
	manufacturerRoute := app.Group("/manufacturer")

	//return all manufacturers
	manufacturerRoute.Get("/", func(c *fiber.Ctx) error {
		manufacturers, err := controllers.GetManufacturers(db)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(manufacturers)
	})

	//create a Manufacturer
	manufacturerRoute.Post("/", func(c *fiber.Ctx) error {
		manufacturer := dto.ManufacturerPost{}

		if err := c.BodyParser(&manufacturer); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := controllers.CreateManufacturer(manufacturer, db); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.SendStatus(201)
	})
}
