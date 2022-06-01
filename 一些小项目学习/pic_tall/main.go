package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"github.com/EdlinOrg/prominentcolor"
)

func main() {

	// Step 1: Load the image
	img, err := loadImage("yes.jpg")
	if err != nil {
		log.Fatal("Failed to load image", err)
	}

	// Step 2: Process it
	colours, err := prominentcolor.Kmeans(img)
	if err != nil {
		log.Fatal("Failed to process image", err)
	}

	fmt.Println("Dominant colours:")
	for _, colour := range colours {
		fmt.Println("#" + colour.AsString())
	}
}

func loadImage(fileInput string) (image.Image, error) {
	f, err := os.Open(fileInput)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	return img, err
}
