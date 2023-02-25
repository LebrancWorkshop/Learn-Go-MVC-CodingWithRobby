package main

import (
	_ "fmt"
	"lebrancconvas/gomvc-robby/initializers"
	"os"

	"github.com/gofiber/fiber/v2"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDatabase()
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":" + os.Getenv("PORT"))
}