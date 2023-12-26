package ffmpeglib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// MP4toH264Converter represents the MP4 to H.264 converter.
type MP4toH264Converter struct {
	InputFile string
}

// NewConverter creates a new MP4toH264Converter instance.
func NewConverter(inputFile string) *MP4toH264Converter {
	return &MP4toH264Converter{
		InputFile: inputFile,
	}
}

// Convert performs the MP4 to H.264 conversion and returns the output file name.
func (c *MP4toH264Converter) Convert() (string, error) {
	// Check if FFmpeg is available
	if err := exec.Command("ffmpeg", "-version").Run(); err != nil {
		return "", fmt.Errorf("FFmpeg is not installed. Please install FFmpeg before using this library.")
	}

	// Generate output file name based on the input file name
	outputFileName := generateOutputFileName(c.InputFile)

	// Run FFmpeg command for conversion
	cmd := exec.Command(
		"/usr/bin/ffmpeg",
		"-i", c.InputFile,
		"-c:v", "libx264", // H.264 codec
		"-preset", "fast", // Choose a preset for encoding speed vs compression ratio
		"-crf", "23", // Constant Rate Factor (0-51, lower is better quality)
		"-c:a", "aac", // AAC audio codec
		"-strict", "experimental", // Needed for older FFmpeg versions
		"-b:a", "128k", // Set audio bitrate (adjust as needed)
		"-threads", "4", // Adjust the number based on your CPU core count
		outputFileName,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("Error converting MP4 to H.264: %v", err)
	}

	return outputFileName, nil
}

// generateOutputFileName generates the output file name with ".mp4" extension based on the input file name.
func generateOutputFileName(inputFileName string) string {
	baseName := filepath.Base(inputFileName)
	ext := filepath.Ext(baseName)
	nameWithoutExt := strings.TrimSuffix(baseName, ext)
	return fmt.Sprintf("%s.mp4", nameWithoutExt)
}
