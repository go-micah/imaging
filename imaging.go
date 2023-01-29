// Package imaging exports a few tools for image processing
package imaging

import (
	"fmt"
	"image"
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
	pointDistance := calculateDistance(400, 400, float64(centerX), float64(centerY))

	fmt.Println(maxDistance)
	fmt.Println(pointDistance)

	return img
}

// calculateDistance calculates the diatnce to the center of the image based on an x,y coordiante
func calculateDistance(x, y, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow((x-x2), 2) + math.Pow((y-y2), 2))
}
