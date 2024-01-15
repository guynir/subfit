package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"subfit/formatutils"
	"subfit/subtitles"
)

func main() {
	inputFile := flag.String("input", "", "Input file")
	outputFile := flag.String("output", "", "Output file")
	adjustOffset := flag.Float64("adjust", 0, "Adjustment offset")
	autoOverwrite := flag.Bool("y", false, "Auto overwrite")
	flag.Parse()

	if formatutils.IsEmpty(*inputFile) {
		fmt.Println("Missing input file.")
		return
	}

	if formatutils.IsEmpty(*outputFile) {
		outputFile = inputFile
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

	if _, err := os.Stat(*outputFile); err == nil && !*autoOverwrite {
		fmt.Printf("Are you sure you want to overwrite '%s' ? ", *outputFile)
		var b []byte = []byte{'n'}

		_, err := os.Stdin.Read(b)
		if err != nil {
			fmt.Println("Aborting")
			return
		}
		str := strings.ToLower(string(b))
		if str == "n" {
			fmt.Println("Aborting")
			return
		}
	}

	err = srt.SaveTo(*outputFile)
	if err != nil {
		fmt.Printf("Failed to write SRT file: %s\n", err.Error())
		return
	}

	fmt.Println("Successfully processed SRT file.")
}
