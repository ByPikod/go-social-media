package handlers

import (
	"errors"
	"net/http"
	"socialmedia/core"
	"socialmedia/helpers"
	"socialmedia/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CommentPayload struct {
	ID           uint   `json:"id"`
	Type         string `json:"type"` // reply or comment
	Content      string `json:"content"`
	AttachmentID *uint  `json:"attachment"`
}

func Comment(c *fiber.Ctx) error {
	// read payload
	var body CommentPayload
	body.Content = c.FormValue("content")
	body.Type = c.FormValue("type")
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.Status(400).Send([]byte("bad request"))
	}
	body.ID = uint(id)
	// check type
	if body.Type != "reply" && body.Type != "comment" {
		return c.Status(400).Send([]byte("bad request"))
	}

	// get user from context
	user := c.Locals("user").(*models.User)

	// check if post exists
	if body.Type == "comment" {
		var post models.Posts
		post.ID = body.ID
		res := core.DB.Where(&post).First(&post)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return c.Status(404).Send([]byte("post not found"))
			}
			return c.Status(500).Send([]byte("internal server error"))
		}
	} else {
		var reply models.Comment
		reply.ID = body.ID
		res := core.DB.Where(reply).First(&reply)
		if res.Error != nil {
			return c.Status(404).Send([]byte("comment to reply not found"))
		}
	}

	// upload image if exists
	file, err := helpers.UploadFile(c, "attachment")
	if err != nil && !errors.Is(err, helpers.ErrNoFile) {
		return c.Status(http.StatusInternalServerError).Send([]byte("internal server error"))
	}

	// create comment
	comment := models.Comment{
		AuthorID: user.ID,
		PostID:   body.ID,
		PostType: body.Type,
		Content:  body.Content,
	}
	if file != nil {
		comment.AttachmentID = file.ID
	}

	res := core.DB.Create(&comment)
	if res.Error != nil {
		return c.Status(500).Send([]byte("internal server error"))
	}

	return c.JSON(fiber.Map{
		"message": "comment created",
		"id":      comment.ID,
	})
}
