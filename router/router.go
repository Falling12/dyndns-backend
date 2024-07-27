package router

import (
	"github.com/gofiber/fiber/v2"
)

func Configure(router fiber.Router) {
	auth := router.Group("/auth")

	authRouter := AuthRouter{}
	auth.Post("/login", authRouter.handleLogin)
	auth.Post("/add-user", authRouter.handleAddUser)
}
