package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
)

func main() {
	fmt.Println("Image to WebP Optimizer")
	fmt.Println("-----------------------")
	fmt.Println("Description: Converts JPG/PNG to WebP (Quality 75%, Max Width 1600px).")
	fmt.Println("Instruction: You can drag and drop an image file directly into this window or manually enter the full file path.")
	fmt.Print("\nPath: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Clean path from drag-and-drop or manual entry
	path := strings.Trim(strings.TrimSpace(input), "\"'")
	path = strings.ReplaceAll(path, `\ `, " ")

	inInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error: File not found. Please check the path and try again.")
		return
	}
	originalSize := inInfo.Size()

	file, _ := os.Open(path)
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error: Invalid image format. Please use JPG or PNG.")
		return
	}

	// Resize logic (Max 1600px)
	bounds := img.Bounds()
	if bounds.Dx() > 1600 {
		newW := 1600
		newH := (bounds.Dy() * newW) / bounds.Dx()
		dst := image.NewRGBA(image.Rect(0, 0, newW, newH))
		draw.CatmullRom.Scale(dst, dst.Bounds(), img, bounds, draw.Over, nil)
		img = dst
	}

	// Create output path with timestamp
	ts := time.Now().Format("150405")
	outPath := strings.TrimSuffix(path, filepath.Ext(path)) + "_" + ts + ".webp"

	outFile, _ := os.Create(outPath)
	webp.Encode(outFile, img, &webp.Options{Quality: 75})
	outFile.Close()

	// Calculate results
	outInfo, _ := os.Stat(outPath)
	newSize := outInfo.Size()

	savings := 0.0
	if originalSize > 0 {
		savings = float64(originalSize-newSize) / float64(originalSize) * 100
	}

	fmt.Printf("\nFinished: %s\n", filepath.Base(outPath))
	fmt.Printf("Compression: %.1f%% smaller\n", savings)
}
