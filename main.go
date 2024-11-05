package main

import (
	"log"
	"wav-to-flac-converter/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Route to handle WAV-to-FLAC conversion
	app.Post("/convert", handlers.ConvertHandler)

	// Start the server on port 8080
	log.Println("Server is running on http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
