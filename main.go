package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yamash723/go-nes-sprites2png/extractor"
)

var (
	romPath    = flag.String("r", "", "input rom file path")
	outputPath = flag.String("o", "", "export png file name")
	width      = flag.Int("w", 400, "output png file width size")
	height     = flag.Int("h", 100, "output png file height size")
)

func main() {
	flag.Parse()

	if *romPath == "" {
		fmt.Println("input rom file path is empty. you should input file path.")
		os.Exit(1)
	}

	if *outputPath == "" {
		fmt.Println("output file path is empty. you should input file path.")
		os.Exit(1)
	}

	err := extractor.Execute(*romPath, *outputPath, *width, *height)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
