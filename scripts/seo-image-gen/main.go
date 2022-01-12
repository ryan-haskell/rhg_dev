package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"path/filepath"
	"strings"

	"os"
)

// TERMINAL CONST
const RESET = "\033[0m"
const RED = "\033[31m"
const GREEN = "\033[32m"
const YELLOW = "\033[33m"
const CYAN = "\033[36m"

// This script is a modified version of this:
// https://www.golangprograms.com/how-to-add-watermark-or-merge-two-image.html

func main() {

	// Step 1: Load background image
	bgFile, err := os.Open("backgrounds/square.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	bgImage, err := png.Decode(bgFile)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer bgFile.Close()

	os.MkdirAll("../../static/gen/banner", os.ModePerm)
	files, err := os.ReadDir("../../static/images")
	if err != nil {
		log.Fatalf("failed to read folder: %s", err)
	}
	count := 0
	for _, file := range files {
		nameWithExt := file.Name()
		filename := strings.TrimSuffix(nameWithExt, filepath.Ext(nameWithExt))
		if filepath.Ext(nameWithExt) == ".png" {
			count++
			createCombinedFile(bgImage, filename)
		}
	}

	// Print success message
	log.Printf(
		"\n\n %s%d%s banner images generated!\n\n",
		GREEN,
		count,
		RESET,
	)
}

func createCombinedFile(bgImage image.Image, filename string) {
	inputFilepath := fmt.Sprintf("../../static/images/%s.png", filename)
	outputFilepath := fmt.Sprintf("../../static/gen/banner/%s.jpg", filename)
	// Step 3: Load foreground image
	fgFile, err := os.Open(inputFilepath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	fgImage, err := png.Decode(fgFile)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer fgFile.Close()

	// Step 4: Describe image dimensions
	bgBounds := image.Rect(0, 0, 256, 256)

	fgOffset := image.Pt(64, 64)
	fgBounds := image.Rect(0, 0, 128, 128).Add(fgOffset)

	// Step 4: Validate image dimensions
	actualFgBounds := fgImage.Bounds().Size()

	if actualFgBounds.X != 128 || actualFgBounds.Y != 128 {
		log.Fatalf(
			"\n\n  Problem with %s./static/images/%s.png%s :\n\n    Expected size %s128x128%s, found %s%dx%d%s\n\n",
			YELLOW,
			filename,
			RESET,
			GREEN,
			RESET,
			CYAN,
			actualFgBounds.X,
			actualFgBounds.Y,
			RESET,
		)
	}

	// Step 5: Create a composite image
	combinedImage := image.NewRGBA(bgBounds)
	draw.Draw(combinedImage, bgBounds, bgImage, image.Point{}, draw.Src)
	draw.Draw(combinedImage, fgBounds, fgImage, image.Point{}, draw.Over)

	// Step 6: Write it to the filesystem
	combinedFile, err := os.Create(outputFilepath)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(combinedFile, combinedImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer combinedFile.Close()
}
