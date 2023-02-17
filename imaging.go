// Package imaging exports a few tools for image processing
package imaging

import (
	"image"
	"image/color"
	"math"
)

// Steps returns a grayscale image divided into 10 equal steps
// https://varis.com/PDFs/DigitalZoneSystem-Part-1.pdf
func Steps(width, height int) image.Image {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewGray(image.Rectangle{upLeft, lowRight})

	increment := width / 11
	zone0 := 0
	zone1 := zone0 + increment
	zone2 := zone1 + increment
	zone3 := zone2 + increment
	zone4 := zone3 + increment
	zone5 := zone4 + increment
	zone6 := zone5 + increment
	zone7 := zone6 + increment
	zone8 := zone7 + increment
	zone9 := zone8 + increment
	zone10 := zone9 + increment

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			// Zone 0
			if x >= zone0 && x <= zone1 {
				img.Set(x, y, color.RGBA{0, 0, 0, 0xff})
			}

			// Zone 1
			if x >= zone1 && x <= zone2 {
				img.Set(x, y, color.RGBA{33, 33, 33, 0xff})
			}

			// Zone 2
			if x >= zone2 && x <= zone3 {
				img.Set(x, y, color.RGBA{51, 51, 51, 0xff})
			}

			// Zone 3
			if x >= zone3 && x <= zone4 {
				img.Set(x, y, color.RGBA{72, 72, 72, 0xff})
			}

			// Zone 4
			if x >= zone4 && x <= zone5 {
				img.Set(x, y, color.RGBA{94, 94, 94, 0xff})
			}

			// Zone 5
			if x >= zone5 && x <= zone6 {
				img.Set(x, y, color.RGBA{118, 118, 118, 0xff})
			}

			// Zone 6
			if x >= zone6 && x <= zone7 {
				img.Set(x, y, color.RGBA{143, 143, 143, 0xff})
			}

			// Zone 7
			if x >= zone7 && x <= zone8 {
				img.Set(x, y, color.RGBA{169, 169, 169, 0xff})
			}

			// Zone 8
			if x >= zone8 && x <= zone9 {
				img.Set(x, y, color.RGBA{197, 197, 197, 0xff})
			}

			// Zone 9
			if x >= zone9 && x <= zone10 {
				img.Set(x, y, color.RGBA{225, 225, 225, 0xff})
			}

			// Zone 10
			if x >= zone10 && x <= width {
				img.Set(x, y, color.RGBA{255, 255, 255, 0xff})
			}
		}
	}

	return img

}

// Dist returns a "distance image" where each pixel's luminance is based on its distance to the center of the image
func Dist(width, height int) image.Image {

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	centerX := width / 2
	centerY := height / 2

	img := image.NewGray(image.Rectangle{upLeft, lowRight})

	maxDistance := calculateDistance(0, 0, float64(centerX), float64(centerY))
	normalizeFactor := 255 / maxDistance

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
