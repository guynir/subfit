package main

import (
	"flag"
	"fmt"
	"subfit/formatutils"
	"subfit/subtitles"
)

func main() {
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "", "Output file")
	adjustOffset := flag.Float64("adjust", 0, "Adjustment offset")
	flag.Parse()

	if formatutils.IsEmpty(*inputFile) {
		fmt.Println("Missing input file.")
		return
	}

	if formatutils.IsEmpty(*outputFile) {
		fmt.Println("Missing output file.")
		return
	}

	srt, err := subtitles.New(*inputFile)
	if err != nil {
		fmt.Printf("Failed to read input SRT file: %s\n", err.Error())
		return
	}

	if *adjustOffset != 0 {
		fmt.Printf("Adjusting subtitles in %.3f seconds.\n", *adjustOffset)
		srt.Adjust(float32(*adjustOffset))
	}

	err = srt.SaveTo(*outputFile)
	if err != nil {
		fmt.Printf("Failed to write SRT file: %s\n", err.Error())
		return
	}

	fmt.Println("Successfully processed SRT file.")
}
