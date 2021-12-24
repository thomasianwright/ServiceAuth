package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"serviceauth/database"
	"serviceauth/models"
	"time"
)

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	if data["key"] == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Key is required"})
	}

	var Application models.Application
	database.DB.Where("public_key = ?", data["public"]).First(&Application)

	if Application.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Application not found"})
	}

	hashed := bcrypt.CompareHashAndPassword([]byte(Application.ApiKey), []byte(data["key"]))

	if hashed != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid key"})
	}

	var scopes []string

	for _, v := range Application.Scopes {
		scopes = append(scopes, v.Name)
	}

	claims := jwt.MapClaims{
		"id":         Application.Id,
		"name":       Application.Name,
		"public_key": Application.PublicKey,
		"exp":        time.Now().Add(time.Hour * 6).Unix(),
		"iat":        time.Now().Unix(),
		"iss":        "serviceauth",
		"scopes":     scopes,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(os.Getenv("jwtsecret")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"token": t})
}
