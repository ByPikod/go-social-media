package handlers

import "github.com/gofiber/fiber/v2"

func Logout(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.Send([]byte("OK!"))
}
