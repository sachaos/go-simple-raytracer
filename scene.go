package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Scene struct {
	height, width int
	camera *Camera
	sphere *Sphere
}

func (s *Scene) Render()  {
	img := image.NewRGBA(image.Rect(0, 0, s.width, s.height))

	for i := 0; i < s.height; i++ {
		for j := 0; j < s.width; j++ {
			ray := s.camera.GetRay(float64(j) / float64(s.width), float64(i)/ float64(s.height))

			pos := s.sphere.Position(ray)
			if pos > 0 {
				N := ray.PointAt(pos).Sub(s.sphere.center)
				r := uint8(math.Abs(Normalize(N).InnerProduct(Normalize(s.sphere.center.Sub(s.camera.pos)))) * 255)
				img.Set(j, s.height-i, color.RGBA{
					r,
					0,
					0,
					255,
				})
			}
		}
	}
	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
