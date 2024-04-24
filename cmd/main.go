package main

import (
	"fmt"
	"kollectionmanager/m/db"
	"kollectionmanager/m/deployment/migrations"
	"kollectionmanager/m/models"
	"kollectionmanager/m/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadEnvFromPath(envFilePath string) {
	err := godotenv.Load(envFilePath)
	if err != nil {
		err = fmt.Errorf("cant load env from path: %s, err {%+v}", envFilePath, err)
		log.Fatal(err)
	}
}

func main() {
	LoadEnvFromPath("/home/rdkvx/documents/_projects/KollectionManager/.env")
	app := fiber.New()

	newDB, err := db.Connect()
	if err != nil {
		fmt.Println(err)
	}

	if !newDB.Migrator().HasTable(&models.Developer{}) {
		migrations.MigrateIfExists(newDB)
		fmt.Println("migration executed successfully")
	}
	

	routes.ConsoleRoutes(app, newDB)

	/* app.Get("/", func(c *fiber.Ctx) error {
	       return c.SendString("Hello, World 👋!")
	   })

	   app.Get("/hello/:name", func(c *fiber.Ctx) error {
	       name := c.Params("name")
	       return c.SendString(fmt.Sprintf("Hello, %s 👋!", name))
	   }) */

	fmt.Println("Running")

	app.Listen(":3000")
}
