package main

import (
	"dyndns/db"
	"dyndns/router"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	api := app.Group("/api")

	db.Connect()
	router.Configure(api)

	app.Listen(":8000")
}
