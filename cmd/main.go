package main

import (
	"fmt"
	handler "go-imagestorage/internal/common"
	imageHandler "go-imagestorage/internal/image"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env", ".env.example");
  if err != nil {
		log.Fatal(err);
		return;
  }
	
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

	fmt.Println("Server is running on localhost:"+os.Getenv("PORT"));
	server.Listen("localhost:"+os.Getenv("PORT"));
}
