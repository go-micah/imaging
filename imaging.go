// Package imaging exports a few tools for image processing
package imaging

import (
	"image"
	"image/color"
	"math"
)

var zoneColors = []color.RGBA{
	{0, 0, 0, 0xff},
	{33, 33, 33, 0xff},
	{51, 51, 51, 0xff},
	{72, 72, 72, 0xff},
	{94, 94, 94, 0xff},
	{118, 118, 118, 0xff},
	{143, 143, 143, 0xff},
	{169, 169, 169, 0xff},
	{197, 197, 197, 0xff},
	{225, 225, 225, 0xff},
	{255, 255, 255, 0xff},
}

func generateZones(width int) []int {
	zoneCount := len(zoneColors)
	increment := width / (zoneCount)
	zones := make([]int, zoneCount)
	for i := 0; i < zoneCount; i++ {
		zones[i] = (i + 1) * increment
	}
	return zones
}

func colorForZone(position int, zones []int) color.RGBA {
	for i, zone := range zones {
		if position < zone {
			return zoneColors[i]
		}
	}
	return zoneColors[len(zoneColors)-1]
}

// Steps returns a grayscale image divided into 10 equal steps
// https://varis.com/PDFs/DigitalZoneSystem-Part-1.pdf
func Steps(width, height int) image.Image {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewGray(image.Rectangle{upLeft, lowRight})

	zones := generateZones(width)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, colorForZone(x, zones))
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
