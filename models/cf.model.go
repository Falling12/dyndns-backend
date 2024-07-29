package models

type NewCloudAccountRequest struct {
	Email string `json:"email"`
	ApiKey string `json:"apiKey"`
}

type CFZone struct {
	ID string `json:"id"`
	Name string `json:"name"`
}