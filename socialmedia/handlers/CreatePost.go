package handlers

import (
	"errors"
	"socialmedia/core"
	"socialmedia/helpers"
	"socialmedia/models"

	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	file, err := helpers.UploadFile(c, "attachment")
	if err != nil && !errors.Is(err, helpers.ErrNoFile) {
		return c.Send([]byte(err.Error()))
	}
	// return if both file and text not provided
	message := c.FormValue("message")
	if file == nil && message == "" {
		return c.Send([]byte("bad request"))
	}
	// create post record in database
	user := c.Locals("user").(*models.User)
	post := &models.Posts{
		AuthorID: user.ID,
	}
	if file != nil {
		post.AttachmentID = file.ID
	}
	if message != "" {
		post.Content = message
	}
	err = core.DB.Create(&post).Error
	if err != nil {
		return c.Send([]byte("failed to create post"))
	}

	// return id as json
	return c.JSON(fiber.Map{
		"id": post.ID,
	})
}
