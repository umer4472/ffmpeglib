package main

import (
	"fmt"

	"github.com/umer4472/ffmpeglib"
)

func main() {
	inputFilePath := "/home/mslm/Downloads/file_example_MOV_1920_2_2MB.mov"

	converter := ffmpeglib.NewConverter(inputFilePath)

	// First conversion
	outputFile, err := converter.Convert()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf(" conversion completed. Output file: %s\n", outputFile)
}
