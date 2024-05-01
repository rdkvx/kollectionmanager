package routes

import (
	"errors"
	"kollectionmanager/m/controllers"
	"kollectionmanager/m/models/dto"
	"kollectionmanager/m/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func GameRoutes(app *fiber.App, db *gorm.DB) {
	gamesRoute := app.Group("/game")

	//return all games
	gamesRoute.Get("/", func(c *fiber.Ctx) error {
		games, err := controllers.GetGames(c, db)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(games)
	})

	//return game by name
	gamesRoute.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		game, err := controllers.GetGameByName(name, db)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(game)
	})

	// create a game
	gamesRoute.Post("/", func(c *fiber.Ctx) error {
		game := dto.GamePost{}

		if err := c.BodyParser(&game); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := controllers.CreateGame(game, db); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.SendStatus(201)
	})

	// update a game by name
	gamesRoute.Patch("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		game := dto.GamePost{}

		if err := c.BodyParser(&game); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if err := controllers.UpdateGame(name, game, db); err != nil {
			if err == errors.New(utils.FailedTo("find", "game", name)){
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                    "error": err.Error(),
                })
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.SendStatus(201)
	})

	gamesRoute.Delete("/:name", func(c *fiber.Ctx) error{
		name := c.Params("name")

        if err := controllers.SoftDeleteGameByName(name, db); err != nil {
            if err.Error() == "record not found" {
				return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
                    "error": err.Error(),
                })
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
        }

        return c.SendStatus(201)
	})
}
