package util

import (
	"fmt"
	"image"
	"net/http"
	"os"
	"strings"

	"github.com/qeesung/image2ascii/convert"

	_ "image/jpeg"
	_ "image/png"
)

func ImageToASCII(imagePath string, width, height uint) (string, error) {
	if width == 0 || height == 0 {
		return "", fmt.Errorf("dimensions must be positive (got %dx%d)", width, height)
	}

	img, err := loadImage(imagePath)
	if err != nil {
		return "", err
	}

	converter := convert.NewImageConverter()
	options := convert.DefaultOptions
	options.FixedWidth = int(width)
	options.FixedHeight = int(height)
	options.FitScreen = false

	ascii := converter.Image2ASCIIString(img, &options)
	return ascii, nil
}

func loadImage(path string) (image.Image, error) {
	var img image.Image

	if strings.HasPrefix(path, "http") {
		resp, err := http.Get(path)
		if err != nil {
			return nil, fmt.Errorf("download failed: %v", err)
		}
		defer resp.Body.Close()

		img, _, err = image.Decode(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("decode failed: %v", err)
		}
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, fmt.Errorf("open failed: %v", err)
		}
		defer file.Close()

		img, _, err = image.Decode(file)
		if err != nil {
			return nil, fmt.Errorf("decode failed: %v", err)
		}
	}

	return img, nil
}

func safeDimensions(requested int) uint {
	if requested <= 0 {
		return 1
	}
	return uint(requested)
}

func GenerateFallbackASCII(width, height uint) string {
	const asciiChars = "\u2588\u2593\u2592"
	var sb strings.Builder
	for y := uint(0); y < height; y++ {
		for x := uint(0); x < width; x++ {
			sb.WriteByte(asciiChars[len(asciiChars)-1]) // Use brightest character
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}
