package imageHandler

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/gofiber/fiber/v2"
)

func Retrieve(c *fiber.Ctx) error {
	category := c.Params("context")
	filename := c.Params("filename")
	return c.SendFile(fmt.Sprintf("assets/images/%s/%s", category, filename))
}

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	filename := c.FormValue("filename")
	context := c.FormValue("context")

	if filename == "" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Filename is required",
		})
	}

	if context == "" {
		return c.Status(400).JSON(&fiber.Map{
			"message": "Context is required",
		})
	}

	if err != nil {
		return c.Status(400).JSON(&fiber.Map{
			"message": "File not found",
		})
	}

	allowedTypes := []string{"image/jpeg", "image/png"}
	if !contains(allowedTypes, file.Header.Get("Content-Type")) {
		return c.Status(400).JSON(&fiber.Map{
			"message": "File type not allowed. Only jpg and png are allowed",
		})
	}

	fileExtension := filepath.Ext(file.Filename)

	destionation := fmt.Sprintf("assets/images/%s/%s%s", context, filename, fileExtension)

	if _, err := os.Stat(fmt.Sprintf("assets/images/%s", context)); os.IsNotExist(err) {
		os.Mkdir(fmt.Sprintf("assets/images/%s", context), 0755)
	}
	saveErr := c.SaveFile(file, destionation)

	if saveErr != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Failed to save file",
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"message":  "File uploaded successfully",
		"url": "http://localhost:8080/images/" + context + "/" + filename + fileExtension,
	})
}

func contains(items []string, item string) bool {
	for _, s := range items {
		if s == item {
			return true
		}
	}
	return false
}
