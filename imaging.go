// Package imaging exports a few tools for image processing
package imaging

import (
	"fmt"
	"image"
	"image/color"
	"math"
)

// Dist returns a "distance image" where each pixel's luminance is based on its distance to the center of the image
func Dist(width, height int) image.Image {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	centerX := width / 2
	centerY := height / 2

	img := image.NewGray(image.Rectangle{upLeft, lowRight})

	maxDistance := calculateDistance(float64(width), float64(height), float64(centerX), float64(centerY))
	fmt.Println(maxDistance)

	normalizeFactor := (maxDistance - 256) / 256

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			pointDistance := calculateDistance(float64(x), float64(y), float64(centerX), float64(centerY))
			pointDistance = pointDistance * normalizeFactor
			gray := color.RGBA{uint8(pointDistance), uint8(pointDistance), uint8(pointDistance), 0xff}
			img.Set(x, y, gray)
		}
	}

	return img
}

// calculateDistance calculates the diatnce to the center of the image based on an x,y coordiante
func calculateDistance(x, y, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow((x-x2), 2) + math.Pow((y-y2), 2))
}
