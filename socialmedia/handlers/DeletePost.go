package handlers

import (
	"errors"
	"socialmedia/core"
	"socialmedia/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type DeletePostStruct struct {
	ID uint `json:"id"`
}

func DeletePost(c *fiber.Ctx) error {
	// read payload
	var body DeletePostStruct
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).Send([]byte("bad request"))
	}

	// get user from context
	user := c.Locals("user").(*models.User)

	// delete post with author id and post id
	var deleted models.Posts
	deleted.ID = body.ID
	deleted.AuthorID = user.ID
	res := core.DB.Delete(&deleted)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.Status(404).Send([]byte("post not found"))
		}
		return c.Status(500).Send([]byte("internal server error"))
	}

	if res.RowsAffected == 0 {
		return c.Status(400).Send([]byte("post not found"))
	}

	return c.Send([]byte("OK!"))
}
