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

	// Step 1: Load background gradient JPEG
	bgImage := loadBackgroundImage()

	// Step 2: Make sure the output folder exists
	ensureOuputFolderExists()

	// Step 3: Create a new file for each PNG image
	count := createImagesWithBackground(bgImage)

	// Step 4: Print out a success message
	printSuccessMessage(count)
}

func loadBackgroundImage() image.Image {
	bgFile, err := os.Open("backgrounds/square.png")
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	bgImage, err := png.Decode(bgFile)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer bgFile.Close()

	return bgImage
}

func ensureOuputFolderExists() {
	os.MkdirAll("../../static/gen/seo", os.ModePerm)
}

func createImagesWithBackground(bgImage image.Image) int {

	// Read all files in the "static/images" folder
	files, err := os.ReadDir("../../static/images")
	if err != nil {
		log.Fatalf("failed to read folder: %s", err)
	}

	// Will track how many files were generated
	count := 0

	for _, file := range files {

		// Get filename, without extension
		nameWithExt := file.Name()
		filename := strings.TrimSuffix(nameWithExt, filepath.Ext(nameWithExt))

		// Ignore non-PNG files
		if filepath.Ext(nameWithExt) == ".png" {
			count++
			createCombinedFile(bgImage, filename)
		}
	}

	return count
}

func printSuccessMessage(count int) {
	log.Printf(
		"\n\n %s%d%s SEO images generated!\n\n",
		GREEN,
		count,
		RESET,
	)
}

// Creates a single file
func createCombinedFile(bgImage image.Image, filename string) {
	// Define input/output paths
	inputFilepath := fmt.Sprintf("../../static/images/%s.png", filename)
	outputFilepath := fmt.Sprintf("../../static/gen/seo/%s.jpg", filename)

	// Load foreground PNG image
	fgFile, err := os.Open(inputFilepath)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	fgImage, err := png.Decode(fgFile)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer fgFile.Close()

	// Describe expected image dimensions
	bgBounds := image.Rect(0, 0, 256, 256)

	fgOffset := image.Pt(64, 64)
	fgBounds := image.Rect(0, 0, 128, 128).Add(fgOffset)

	// Validate against actual image dimensions
	actualFgBounds := fgImage.Bounds().Size()

	// Exit with helpful error if a PNG is the wrong size
	if actualFgBounds.X != 128 || actualFgBounds.Y != 128 {
		printInvalidImageError(filename, actualFgBounds)
	}

	// Create a composite image
	combinedImage := image.NewRGBA(bgBounds)
	draw.Draw(combinedImage, bgBounds, bgImage, image.Point{}, draw.Src)
	draw.Draw(combinedImage, fgBounds, fgImage, image.Point{}, draw.Over)

	// Write output to the filesystem
	combinedFile, err := os.Create(outputFilepath)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(combinedFile, combinedImage, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer combinedFile.Close()
}

func printInvalidImageError(filename string, actualBounds image.Point) {
	log.Fatalf(
		"\n\n  Problem with %s./static/images/%s.png%s :\n\n    Expected size %s128x128%s, found %s%dx%d%s\n\n",
		YELLOW,
		filename,
		RESET,
		GREEN,
		RESET,
		CYAN,
		actualBounds.X,
		actualBounds.Y,
		RESET,
	)
}
