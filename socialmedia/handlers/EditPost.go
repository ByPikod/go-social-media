package handlers

import (
	"errors"
	"socialmedia/core"
	"socialmedia/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type EditPostPayload struct {
	ID      uint   `json:"id"`
	Content string `json:"content"`
}

func EditPost(c *fiber.Ctx) error {
	// get payload
	var payload EditPostPayload
	err := c.BodyParser(&payload)
	if err != nil {
		return c.Send([]byte("bad request"))
	}
	// retrieve post
	user := c.Locals("user").(*models.User)
	post := models.Posts{}
	err = core.DB.Where("id=?", payload.ID).First(&post).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Send([]byte("post not found"))
		}
		return c.Send([]byte("internal server error"))
	}
	// check if user is author
	if post.AuthorID != user.ID {
		return c.Send([]byte("unauthorized"))
	}
	// update post
	post.Content = payload.Content
	err = core.DB.Save(&post).Error
	if err != nil {
		return c.Send([]byte("failed to update post"))
	}
	// return success
	return c.Send([]byte("success"))
}
