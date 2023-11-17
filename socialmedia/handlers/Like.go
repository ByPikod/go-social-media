package handlers

import (
	"errors"
	"socialmedia/core"
	"socialmedia/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type LikePayload struct {
	ID   uint   `json:"id"`
	Type string `json:"type"` // reply or comment
}

func Like(c *fiber.Ctx) error {
	// read payload
	var body LikePayload
	err := c.BodyParser(&body)
	if err != nil {
		return c.Status(400).Send([]byte("bad request"))
	}

	// check type
	if body.Type != "post" && body.Type != "comment" {
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

	// create like
	like := models.Like{
		AuthorID: user.ID,
		PostID:   body.ID,
		PostType: body.Type,
	}
	res := core.DB.Create(&like)
	if res.Error == nil {
		return c.Send([]byte("liked"))
	}

	if !errors.Is(res.Error, gorm.ErrDuplicatedKey) {
		// should unlike
		return c.Status(500).Send([]byte("internal server error"))
	}

	// delete like (unlike)
	res = core.DB.Where(&like).Delete(&like)
	if res.Error != nil {
		return c.Status(500).Send([]byte("internal server error"))
	}

	return c.Status(404).Send([]byte("unliked"))
}
