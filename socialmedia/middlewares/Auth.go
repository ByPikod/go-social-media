package middlewares

import (
	"errors"
	"socialmedia/core"
	"socialmedia/helpers"
	"socialmedia/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Auth(c *fiber.Ctx) error {
	// verify cookie exist
	cookie := c.Cookies("token")
	if cookie == "" {
		return c.Status(401).Send([]byte("unauthorized"))
	}
	// verify token
	verified, err := helpers.VerifyToken(cookie)
	if err != nil {
		return c.Status(500).Send([]byte("internal server error"))
	}
	if verified == "" {
		return c.Status(401).Send([]byte("unauthorized"))
	}
	// retrieve user from token
	user := models.User{
		Username: verified,
	}
	err = core.DB.Model(&models.User{}).Where(&user).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(401).Send([]byte("account suspended"))
		}
		return c.Status(500).Send([]byte("internal server error"))
	}
	c.Locals("user", &user)
	return c.Next()
}
