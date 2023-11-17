package handlers

import (
	"socialmedia/core"
	"socialmedia/models"

	"github.com/gofiber/fiber/v2"
)

type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *fiber.Ctx) error {
	// parse body
	var body RegisterPayload
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).Send([]byte("bad request"))
	}
	// check form data
	if body.Username == "" || body.Password == "" {
		return c.Status(400).Send([]byte("bad request"))
	}
	// check if user exists
	var exists bool
	err := core.DB.Model(&models.User{}).
		Select("count(*) > 0").
		Where(&models.User{
			Username: body.Username,
		}).
		Find(&exists).
		Error
	if err != nil {
		return c.Status(500).Send([]byte("internal server error"))
	}
	if exists {
		return c.Status(400).Send([]byte("user already exists"))
	}
	// create user
	user := models.User{
		Username: body.Username,
		Password: body.Password,
	}
	err = core.DB.Create(&user).Error
	if err != nil {
		return c.Status(500).Send([]byte("internal server error"))
	}
	return c.Send([]byte("successfully created"))
}
