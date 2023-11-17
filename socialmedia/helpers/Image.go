package helpers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"socialmedia/core"
	"socialmedia/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

var ErrNoFile error = errors.New("no file")

// allowed formats
var formats = map[string]bool{
	"image/jpeg": true,
	"image/png":  true,
}

func VerifyImage(file *multipart.FileHeader, maxSizeMB int64) error {
	// check format
	_, ok := formats[file.Header.Get("Content-Type")]
	if !ok {
		return errors.New("invalid file format")
	}
	// check size
	if file.Size > (maxSizeMB * 1024 * 1024) {
		return errors.New("file too large")
	}
	return nil
}

func UploadFile(c *fiber.Ctx, fileFormName string) (*models.Files, error) {
	// get file from form
	file, err := c.FormFile(fileFormName)
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		return nil, ErrNoFile
	}
	// check file
	err = VerifyImage(file, 8)
	if err != nil {
		return nil, err
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
		return nil, errors.New("failed to save file")
	}
	// create file record in database
	fileRecord := &models.Files{
		FilePath: filename,
	}
	err = core.DB.Create(fileRecord).Error
	if err != nil {
		return nil, errors.New("failed to save file into database")
	}

	// successful
	return fileRecord, nil
}
