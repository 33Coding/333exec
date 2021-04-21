package main

import (
	"333exec/Runner"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	runner := Runner.Runner{}

	app.Get("/get/*", runner.Handler)

	app.Listen(":80")
}
