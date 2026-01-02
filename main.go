package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/chai2010/webp"
	"golang.org/x/image/draw"
)

func main() {
	var inputPath string
	var qualityStr string
	quality := 75

	// Case A: Running with arguments (Docker or CLI)
	if len(os.Args) >= 2 {
		inputPath = os.Args[1]
		if len(os.Args) > 2 {
			qualityStr = os.Args[2]
		}
	} else {
		// Case B: Interactive Mode (Local execution)
		fmt.Println("Image to WebP Optimizer")
		fmt.Println("-----------------------")
		fmt.Println("Converts JPG/PNG to WebP")
		fmt.Print("\nPath: ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			inputPath = scanner.Text()
		}

		fmt.Print("Quality (Default 75): ")
		if scanner.Scan() {
			qualityStr = scanner.Text()
		}
	}

	// Clean Path
	inputPath = strings.Trim(strings.TrimSpace(inputPath), "\"'")
	inputPath = strings.ReplaceAll(inputPath, `\ `, " ")

	if inputPath == "" {
		fmt.Println("Error: No path provided.")
		return
	}

	// Parse Quality
	if qualityStr != "" {
		if q, err := strconv.Atoi(strings.TrimSpace(qualityStr)); err == nil {
			quality = q
		}
	}

	fileInfo, err := os.Stat(inputPath)
	if err != nil {
		fmt.Printf("Error: Path '%s' not found.\n", inputPath)
		return
	}

	if fileInfo.IsDir() {
		fmt.Printf("\nBatch processing folder: %s (Quality: %d%%)\n", inputPath, quality)
		files, _ := os.ReadDir(inputPath)
		for _, f := range files {
			ext := strings.ToLower(filepath.Ext(f.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				processImage(filepath.Join(inputPath, f.Name()), quality)
			}
		}
	} else {
		processImage(inputPath, quality)
	}
}

func processImage(path string, quality int) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error opening %s: %v\n", path, err)
		return
	}
	defer file.Close()

	inInfo, _ := file.Stat()
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Printf("Error decoding %s: Use JPG or PNG\n", path)
		return
	}

	bounds := img.Bounds()
	if bounds.Dx() > 1600 {
		newW := 1600
		newH := (bounds.Dy() * newW) / bounds.Dx()
		dst := image.NewRGBA(image.Rect(0, 0, newW, newH))
		draw.CatmullRom.Scale(dst, dst.Bounds(), img, bounds, draw.Over, nil)
		img = dst
	}

	ts := time.Now().Format("150405")
	outPath := strings.TrimSuffix(path, filepath.Ext(path)) + "_" + ts + ".webp"

	outFile, err := os.Create(outPath)
	if err != nil {
		fmt.Printf("Error creating file %s: %v\n", outPath, err)
		return
	}
	defer outFile.Close()

	webp.Encode(outFile, img, &webp.Options{Quality: float32(quality)})

	outInfo, _ := os.Stat(outPath)
	savings := float64(inInfo.Size()-outInfo.Size()) / float64(inInfo.Size()) * 100
	fmt.Printf("✔ %s | Reduced: %.1f%% (Quality: %d)\n", filepath.Base(outPath), savings, quality)
}
