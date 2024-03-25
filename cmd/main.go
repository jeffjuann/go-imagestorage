package main

import (
	"fmt"
	handler "go-imagestorage/internal/common"
	imageHandler "go-imagestorage/internal/image"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := fiber.New()
	server.Static("/", "./public")

	server.Get("/", handler.Info)
	imageGroup := server.Group("/images")
	imageGroup.Get("/:context/:filename", imageHandler.Retrieve)
	imageGroup.Post("/upload", imageHandler.Upload)

	port := 3000

	fmt.Println("Server is running on http://localhost:" + strconv.Itoa(port))
	server.Listen("localhost:" + strconv.Itoa(port))
}
