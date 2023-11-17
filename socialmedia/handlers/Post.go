package handlers

import (
	"errors"
	"socialmedia/core"
	"socialmedia/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Post(c *fiber.Ctx) error {
	idstr := c.Params("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		return c.SendString("bad request")
	}
	// get post with id
	post := models.Posts{}
	post.ID = uint(id)
	res := core.DB.Where(&post).Preload("Author").First(&post)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return c.SendString("post not found")
		}
		return c.SendString("internal server error")
	}

	ret := struct {
		PostID     uint    `json:"post_id"`
		Content    string  `json:"content"`
		Author     string  `json:"author"`
		Attachment *string `json:"attachment"`
	}{
		PostID:  post.ID,
		Content: post.Content,
		Author:  post.Author.Username,
	}

	// get attachment if exists
	if post.AttachmentID != 0 {
		attachment := models.Files{}
		attachment.ID = post.AttachmentID
		res = core.DB.Where(&attachment).First(&attachment)
		if res.Error == nil {
			str := "cdn.example.com/" + attachment.FilePath
			ret.Attachment = &str
		}
	}

	return c.JSON(ret)
}
