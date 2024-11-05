 # WAV to FLAC Converter

## Overview

This project is a Go service that converts WAV audio streams to FLAC format in real-time using WebSockets.

## Setup Instructions

1. Ensure you have [Go](https://golang.org/dl/) installed.
2. Clone this repository or download the project files.
3. Navigate to the project directory and run:
   ```bash
   go mod tidy
   go run main.go
Install dependencies:
go mod tidy
Build the project:
go build -o converter main.go
Run the service:
./converter

Convert Audio
POST /converter

Request

Body: Binary WAV file
Response

Content-Type: audio/flac
Body: Converted FLAC file

Example using curl
curl -X POST -H "Content-Type: audio/wav" --data-binary @path/to/your.wav http://localhost:8080/converter -o output.flac

Running Tests
To run the tests, use the command:
go test ./...
