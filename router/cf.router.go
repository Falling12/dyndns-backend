package router

import (
	"dyndns/controllers"
	"dyndns/models"

	"github.com/gofiber/fiber/v2"
)

type CFRouter struct{}
var cfController = controllers.CFController{}

func (r *CFRouter) handleGetZones(c *fiber.Ctx) error {
	cloudflareId := c.Params("id")

	zones, err := cfController.HandleGetZones(cloudflareId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(zones)
}

func (r *CFRouter) handleNewCF(c *fiber.Ctx) error {
	req := new(models.NewCloudAccountRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cf, err := cfController.HandleNewCF(*req)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(cf)
}