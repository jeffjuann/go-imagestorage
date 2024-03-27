package main

import (
	"fmt"
	handler "go-imagestorage/internal/common"
	imageHandler "go-imagestorage/internal/image"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	server := fiber.New()
	server.Static("/", "./public")
	server.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
	}))

	server.Get("/", handler.Info)
	imageGroup := server.Group("/images")
	imageGroup.Get("/:context/:filename", imageHandler.Retrieve)
	imageGroup.Post("/upload", imageHandler.Upload)

	port := 8080

	fmt.Println("Server is running on http://localhost:" + strconv.Itoa(port))
	server.Listen("localhost:" + strconv.Itoa(port))
}
