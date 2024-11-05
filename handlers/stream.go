package handlers

import (
	"fmt"
	"wav-to-flac-converter/ffmpeg"

	"github.com/gofiber/fiber/v2"
)

// ConvertHandler processes the audio conversion from WAV to FLAC
func ConvertHandler(c *fiber.Ctx) error {
	// Verify content type
	if c.Get("Content-Type") != "audio/wav" {
		return c.Status(fiber.StatusBadRequest).SendString("Content-Type must be audio/wav")
	}

	// Read the body containing WAV data
	wavData := c.Body()
	if len(wavData) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("No WAV data received")
	}

	// Convert WAV data to FLAC
	flacData, err := ffmpeg.ConvertToFlac(wavData)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(fmt.Sprintf("Conversion error: %v", err))
	}

	// Send the FLAC data as response
	c.Set("Content-Type", "audio/flac")
	return c.Send(flacData)
}
