package controllers

import (
	"dyndns/db"
	"dyndns/models"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct{}

func (c *AuthController) HandleLogin(data models.LoginRequest) (string, error) {
    user, err := db.DB.User.FindFirst(db.User.Name.Equals(data.Username)).Exec(db.Ctx)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("Invalid username or password")
	}

	if user.Password != data.Password {
		return "", errors.New("Invalid username or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
		"username": user.Name,
		"email": user.Email,
	})

	str, err := token.SignedString([]byte("secret"))

	if err != nil {
		return "", err
	}

	return str, nil
}