package routes

import (
	"fmt"
	"kollectionmanager/m/controllers"
	"kollectionmanager/m/models"
	"kollectionmanager/m/models/dto"
	"kollectionmanager/m/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func DeveloperRoutes(app *fiber.App, db *gorm.DB) {
	developersRoute := app.Group("/developer")

	//return all developers
	developersRoute.Get("/", func(c *fiber.Ctx) error {
		developers, err := controllers.GetDevelopers(db)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(developers)
	})

	//return developer by name
	developersRoute.Get("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		developer, err := controllers.GetDeveloperByName(name, db)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(developer)
	})

	//create a developer
	developersRoute.Post("/", func(c *fiber.Ctx) error {
		developer := models.Developer{}

		if err := c.BodyParser(&developer); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
		}

		fmt.Println(developer)

		err := controllers.CreateDeveloper(developer, db)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": err.Error(),
            })
		}

		return c.SendStatus(201)
	})

	//update a developer by name
	developersRoute.Patch("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		developer := dto.Developer{}

		if err := c.BodyParser(&developer); err != nil {
			return err
		}

		err := controllers.UpdateDeveloperByName(name, developer, db)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		fmt.Println(utils.DevUpdatedSuccess)

		return c.SendStatus(200)
	})

	//delete a developer by name
	developersRoute.Delete("/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		if err := controllers.SoftDeleteDeveloperByName(name, db); err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		fmt.Println(utils.DevDeletedSuccess)

		return c.SendStatus(200)

	})
}
