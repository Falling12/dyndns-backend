package controllers

import (
	"dyndns/db"
	"dyndns/models"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type CFController struct{}

func (c *CFController) HandleNewCF(data models.NewCloudAccountRequest) (*db.CloudflareModel, error) {
	cf, err := db.DB.Cloudflare.CreateOne(
		db.Cloudflare.Email.Set(data.Email),
		db.Cloudflare.APIKey.Set(data.ApiKey),
	).Exec(db.Ctx)

	if err != nil {
		return nil, err
	}

	return cf, nil
}

func (c *CFController) HandleGetZones(cloudflareId string) ([]models.CFZone, error) {
	cf, err := db.DB.Cloudflare.FindFirst(db.Cloudflare.ID.Equals(cloudflareId)).Exec(db.Ctx)

	if err != nil {
		return nil, err
	}

	if cf == nil {
		return nil, fiber.NewError(fasthttp.StatusNotFound, "Cloudflare account not found")
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("https://api.cloudflare.com/client/v4/zones")
	req.Header.SetMethod("GET")
	req.Header.Set("X-Auth-Key", cf.APIKey)
	req.Header.Set("X-Auth-Email", cf.Email)

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}

	if err := client.Do(req, resp); err != nil {
		return nil, err
	}

	if resp.StatusCode() != fasthttp.StatusOK {
		fmt.Println(resp)
		return nil, fiber.NewError(fasthttp.StatusInternalServerError, "Failed to fetch zones")
	}

	var result map[string]interface{}
	if err := json.Unmarshal(resp.Body(), &result); err != nil {
		return nil, err
	}

	zones := make([]models.CFZone, 0)

	for _, zone := range result["result"].([]interface{}) {
		z := zone.(map[string]interface{})
		zones = append(zones, models.CFZone{
			ID:   z["id"].(string),
			Name: z["name"].(string),
		})
	}

	return zones, nil
}