package handlers

import (
	"fmt"
	"path/filepath"
	"socialmedia/core"
	"socialmedia/helpers"
	"socialmedia/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ProfileUpdate(c *fiber.Ctx) error {
	file, err := c.FormFile("picture")
	if err != nil {
		return c.Send([]byte("failed to process file"))
	}
	// check file
	err = helpers.VerifyImage(file, 8)
	if err != nil {
		return c.Send([]byte(err.Error()))
	}
	// move file to ./uploads
	user := c.Locals("user").(*models.User)
	filename := fmt.Sprintf(
		"%s-%s.%s",
		user.Username,
		time.Now().Format("2006-01-02-15-04-05"),
		filepath.Ext(file.Filename),
	)
	err = c.SaveFile(file, "./uploads/"+filename)
	if err != nil {
		return c.Send([]byte("failed to save file"))
	}
	// create file record in database
	fileRecord := &models.Files{
		FilePath: filename,
	}
	err = core.DB.Create(fileRecord).Error
	if err != nil {
		return c.Send([]byte("failed to save file into database"))
	}
	// update user
	user.PFP = fileRecord.ID
	err = core.DB.Save(user).Error
	if err != nil {
		return c.Send([]byte("failed to update file field for user in database"))
	}

	return c.Send([]byte("OK!"))
}
