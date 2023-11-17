package handlers

import (
	"errors"
	"socialmedia/core"
	"socialmedia/helpers"
	"socialmedia/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	// parse body
	var body LoginPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).Send([]byte("bad request"))
	}
	// check form data
	if body.Username == "" || body.Password == "" {
		return c.Status(400).Send([]byte("bad request"))
	}
	// check if user exists
	var exists models.User
	err := core.DB.Model(&models.User{}).
		Where(&models.User{
			Username: body.Username,
		}).
		Find(&exists).
		Error
	if err != nil {
		// if user does not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(400).Send([]byte("incorrect credentials"))
		}
		// anything else
		return c.Status(500).Send([]byte("internal server error at user check"))
	}
	// check if password is correct
	if exists.Password != body.Password {
		return c.Status(400).Send([]byte("incorrect credentials"))
	}
	// create token
	token, err := helpers.CreateToken(body.Username)
	if err != nil {
		return c.Status(500).Send([]byte("internal server error: " + err.Error()))
	}
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: token,
	})
	// return token
	return c.Send([]byte("OK!"))
}
