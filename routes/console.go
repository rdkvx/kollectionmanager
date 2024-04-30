package routes

import (
	"kollectionmanager/m/controllers"
	"kollectionmanager/m/models/dto"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ConsoleRoutes(app *fiber.App, db *gorm.DB) {
	consolesroute := app.Group("/console")

	//return all consoles
	consolesroute.Get("/", func(c *fiber.Ctx) error {
		consoles, err := controllers.GetConsoles(c, db)
		if err != nil {
            return c.SendStatus(500)
        }

		return c.JSON(consoles)
	})

	//return console by name
	consolesroute.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

        console, err := controllers.GetConsoleByName(name, db)
        if err != nil {
            return c.SendStatus(500)
        }

        return c.JSON(console)
	})

	//create a console
	consolesroute.Post("/", func(c *fiber.Ctx) error {
		console := dto.ConsolePost{}

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

	consolesroute.Patch("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
        console := dto.ConsolePost{}

        if err := c.BodyParser(&console); err != nil {
            return err
        }

        if err := controllers.UpdateConsole(name, console, db); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.SendStatus(201)
	})

	consolesroute.Delete("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

        if err := controllers.DeleteConsole(name, db); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.SendStatus(201)
	})
}