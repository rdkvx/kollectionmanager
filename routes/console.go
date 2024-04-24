package routes

import (
	"kollectionmanager/m/controllers"

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
}