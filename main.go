package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width := 200
	height := 200

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			img.Set(j, i, color.RGBA{255, 255, 255, 255})
		}
	}

	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
