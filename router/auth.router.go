package router

import (
	"github.com/gofiber/fiber/v2"

	"dyndns/controllers"
	"dyndns/db"
	"dyndns/models"
)

type AuthRouter struct{}

var authController = controllers.AuthController{}
func (r *AuthRouter) handleLogin(c *fiber.Ctx) error {
	req := new(models.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := authController.HandleLogin(*req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (r *AuthRouter) handleMe(c *fiber.Ctx) error {
	return c.SendString("Me")
}

func (r *AuthRouter) handleAddUser(c *fiber.Ctx) error {
	user, err := db.DB.User.CreateOne(
		db.User.Name.Set("admin"),
		db.User.Email.Set("senkcsani@gmail.com"),
		db.User.Password.Set("admin"),
	).Exec(db.Ctx)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(user)
}