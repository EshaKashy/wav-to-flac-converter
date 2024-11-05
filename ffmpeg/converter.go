package ffmpeg

import (
	"bytes"
	"os/exec"
)

// ConvertToFlac converts WAV data to FLAC format using ffmpeg
func ConvertToFlac(wavData []byte) ([]byte, error) {
	// Create an ffmpeg command to convert WAV to FLAC
	cmd := exec.Command("ffmpeg", "-i", "pipe:0", "-f", "flac", "pipe:1")

	// Prepare the input and output buffers
	var flacOutput bytes.Buffer
	cmd.Stdin = bytes.NewReader(wavData)
	cmd.Stdout = &flacOutput

	// Run the ffmpeg command
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	// Return the converted FLAC data
	return flacOutput.Bytes(), nil
}
