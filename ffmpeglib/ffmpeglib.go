package ffmpeglib

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ConversionOptions represents the options for MP4 to H264 conversion.
type ConversionOptions struct {
	VideoCodec   string
	AudioBitrate string
	AudioCodec   string
}

// MP4toH264Converter represents the MP4 to H264 converter.
type MP4toH264Converter struct {
	ConversionOptions
	InputFile string
	Preset    string
	CRF       string
	Strict    string
	Threads   string
}

// NewConverter creates a new MP4toH264Converter instance.
func NewConverter(inputFile string, options ConversionOptions) (*MP4toH264Converter, error) {

	if _, err := os.Stat(inputFile); os.IsNotExist(err) {
		return nil, fmt.Errorf("Input file does not exist: %s", inputFile)
	}

	return &MP4toH264Converter{
		ConversionOptions: options,
		InputFile:         inputFile,
		Preset:            "fast",
		CRF:               "23",
		Strict:            "experimental",
		Threads:           "4",
	}, nil
}

// Convert performs the MP4 to H.264 conversion and returns the output file name.
func (c *MP4toH264Converter) Convert() (string, error) {
	ffmpegPath := "/usr/bin/ffmpeg" // Set a default path

	// Check if FFmpeg is available
	if err := exec.Command("ffmpeg", "-version").Run(); err != nil {
		return "", fmt.Errorf("FFmpeg is not installed. Please install FFmpeg before using this library.")
	}

	// Generate output file name based on the input file name
	outputFileName := generateOutputFileName(c.InputFile)

	// Run FFmpeg command for conversion
	cmd := exec.Command(
		ffmpegPath,
		"-i", c.InputFile,
		"-c:v", c.VideoCodec,
		"-preset", c.Preset,
		"-crf", c.CRF,
		"-c:a", c.AudioCodec,
		"-strict", c.Strict,
		"-b:a", c.AudioBitrate,
		"-threads", c.Threads,
		outputFileName,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Error converting MP4 to H.264: %v", err)
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
