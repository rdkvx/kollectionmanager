package routes

import (
	"kollectionmanager/m/controllers"
	"kollectionmanager/m/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ConsoleRoutes(app *fiber.App, db *gorm.DB) {
	consolesroute := app.Group("/consoles")

	consolesroute.Get("/", func(c *fiber.Ctx) error {
		consoles, err := controllers.GetConsoles(c, db)
		if err != nil {
            return c.SendStatus(500)
        }

		return c.JSON(consoles)
	})

	consolesroute.Post("/", func(c *fiber.Ctx) error {
		console := models.Console{}

        if err := c.BodyParser(&console); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        if err := controllers.CreateConsole(console, db); err != nil{
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.SendStatus(201)
	})
}