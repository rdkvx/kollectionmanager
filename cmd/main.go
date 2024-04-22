package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World 👋!")
    })

    app.Get("/hello/:name", func(c *fiber.Ctx) error {
        name := c.Params("name")
        return c.SendString(fmt.Sprintf("Hello, %s 👋!", name))
    })

	fmt.Println("Running")

    app.Listen(":3000")
}
