package ffmpeglib

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// ConversionOptions represents the options for MP4 to H264 conversion.
type ConversionOptions struct {
	InputFile    string
	VideoCodec   string
	AudioBitrate string
}

// MP4toH264Converter represents the MP4 to H264 converter.
type MP4toH264Converter struct {
	ConversionOptions
	Preset     string
	CRF        string
	AudioCodec string
	Strict     string
	Threads    string
}

// NewConverter creates a new MP4toH264Converter instance.
func NewConverter(options ConversionOptions) *MP4toH264Converter {
	return &MP4toH264Converter{
		ConversionOptions: options,
		Preset:            "fast",
		CRF:               "23",
		AudioCodec:        "aac",
		Strict:            "experimental",
		Threads:           "4",
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
