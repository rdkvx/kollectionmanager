package routes

import (
	"kollectionmanager/m/controllers"
	"kollectionmanager/m/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GameRoutes(app *fiber.App, db *gorm.DB){
	gamesRoute := app.Group("/game")

	gamesRoute.Post("/game", func(c *fiber.Ctx) error{
		game := models.Game{}

		if err := c.BodyParser(&game); err!= nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

		if err := controllers.CreateGame(game, db); err != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
		}

		return c.SendStatus(201)
	})
}