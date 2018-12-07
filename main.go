package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func main() {
	height := 200
	width := 200

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	lookFrom := NewVector(-2, 0, 0)
	lookAt := NewVector(2, 0, 0)
	vup := NewVector(0, 0, 1)

	c := NewCamera(lookFrom, lookAt, vup, 0.5, 1.0)
	sphere := &Sphere{NewVector(0, 0.0, 0.2), 0.2}
	look := lookAt.Sub(lookFrom)
	fmt.Println(c)
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			ray := c.GetRay(float64(j) / float64(width), float64(i)/ float64(height))

			pos := sphere.Position(ray)
			if pos > 0 {
				N := ray.PointAt(pos).Sub(sphere.center)
				s := uint8(math.Abs(Normalize(N).InnerProduct(Normalize(sphere.center.Sub(c.pos)))) * 255)
				print(s)
				print(" ")
				img.Set(j, height-i, color.RGBA{
					s,
					0,
					0,
					255,
				})
			} else {
				look := math.Abs(Normalize(ray.Dir).InnerProduct(Normalize(look)))
				unconcern := uint8(math.Pow(look, 32.0) * 255.0)
				print(unconcern)
				print(" ")
				img.Set(j, height-i, color.RGBA{
					unconcern,
					unconcern,
					unconcern,
					255,
				})
			}
		}
		println()
	}
	println(look.String())

	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
