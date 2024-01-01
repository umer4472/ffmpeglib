package main

import (
	"fmt"

	"github.com/umer4472/ffmpeglib"
)

func main() {
	options := ffmpeglib.ConversionOptions{
		InputFile:    "/home/mslm/Downloads/file_example_MOV_1920_2_2MB.mov",
		VideoCodec:   "libx264",
		AudioBitrate: "192k",
	}

	converter := ffmpeglib.NewConverter(options)

	outputFile, err := converter.Convert()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Conversion completed. Output file: %s\n", outputFile)
}
