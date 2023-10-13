package controllers

import (
	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
	"go_starter/config"
	"go_starter/models"
	"strconv"
	"time"
)

func GenerateTokenWeb(user models.User) (string, error) {
	// Extract the credentials from the request body
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"id":       strconv.Itoa(int(user.ID)),
		"username": user.Username,
		"name":     user.Name,
		"exp":      time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SecretWeb))
	if err != nil {
		return "", err
	}
	return t, nil
}

func GenerateTokenApi(partner models.Partner) (string, error) {
	// Extract the credentials from the request body
	day := time.Hour * 24
	// Create the JWT claims, which includes the user ID and expiry time
	claims := jtoken.MapClaims{
		"id":       strconv.Itoa(int(partner.ID)),
		"username": partner.Username,
		"name":     partner.Name,
		"exp":      time.Now().Add(day * 1).Unix(),
	}
	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.SecretApi))
	if err != nil {
		return "", err
	}
	return t, nil
}

func Protected(c *fiber.Ctx) error {
	// Get the user from the context and return it
	user := c.Locals("user").(*jtoken.Token)
	claims := user.Claims.(jtoken.MapClaims)
	email := claims["username"].(string)
	return c.SendString("Welcome 👋" + email)
}
