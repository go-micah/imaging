package main

import (
	"image/png"
	"os"

	"github.com/go-micah/imaging"
)

func main() {

	img := imaging.Dist(512, 512)

	f, _ := os.Create("./output/dist.png")

	png.Encode(f, img)

}
