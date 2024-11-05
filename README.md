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

# WAV to FLAC Converter Service

## Overview
This service allows you to upload WAV audio files, which are then converted to FLAC format.

## Prerequisites
- Go 1.23 or later
- FFmpeg installed and added to your system PATH

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/wav-to-flac-converter.git
   cd wav-to-flac-converter