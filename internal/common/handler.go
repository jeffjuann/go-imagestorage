package handler

import (
	"github.com/gofiber/fiber/v2"
)

func Info(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"message": "Go-SoundStorage is running",
	});
}