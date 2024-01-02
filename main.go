package main

import (
	"fmt"
	"log"

	"github.com/umer4472/ffmpeglib"
)

func main() {
	options := ffmpeglib.ConversionOptions{
		VideoCodec:   "libx264",
		AudioBitrate: "192k",
		AudioCodec:   "aac",
		Preset:       "fast", // compression ratio
		CRF:          "23",   //video quality
	}

	inputFile := "/home/mslm/Downloads/file_example_MOV_1920_2_2MB.mov"

	converter, err := ffmpeglib.NewConverter(inputFile, options)

	if err != nil {
		log.Fatal(err)
	}

	outputFile, err := converter.Convert()

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Conversion completed. Output file: %s\n", outputFile)
}
