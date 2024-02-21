package main

import (
	"fetchAPI_gofiber/src/config"
	"fetchAPI_gofiber/src/helper"
	"fetchAPI_gofiber/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.InitDB()
	helper.Migration()
	routes.Router(app)

	app.Listen(":8080")
}
