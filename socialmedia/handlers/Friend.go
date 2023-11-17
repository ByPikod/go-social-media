package handlers

import (
	"errors"
	"socialmedia/core"
	"socialmedia/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AddFriendPayload struct {
	Username string `json:"username"`
}

func Friend(c *fiber.Ctx) error {
	// parse payload
	var payload AddFriendPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(400).Send([]byte("bad request"))
	}

	// get user
	user := c.Locals("user").(*models.User)

	// check if user to add exists
	friend := models.User{
		Username: payload.Username,
	}
	err := core.DB.Where(&friend).First(&friend).Error
	if err != nil {
		return c.Status(400).Send([]byte("user not found"))
	}

	// create friendship
	friendship := models.Friends{
		Inviting:  user.ID,
		Accepting: friend.ID,
	}
	res := core.DB.Create(&friendship).Error
	if res == nil {
		return c.SendString("friend added.")
	}

	if !errors.Is(res, gorm.ErrDuplicatedKey) {
		return c.Status(500).Send([]byte("internal server error"))
	}

	// unfriend
	err = core.DB.Where(&friendship).Delete(&friendship).Error
	if err != nil {
		return c.Status(500).Send([]byte("internal server error"))
	}

	return c.SendString("friend removed.")
}
