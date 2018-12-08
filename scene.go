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
	spheres []*Sphere
}

func (s *Scene) Render()  {
	img := image.NewRGBA(image.Rect(0, 0, s.width, s.height))

	for i := 0; i < s.height; i++ {
		for j := 0; j < s.width; j++ {
			ray := s.camera.GetRay(float64(j) / float64(s.width), float64(i)/ float64(s.height))

			minIndex := -1
			minPosition := math.MaxFloat64
			for i, sphere := range s.spheres {
				position := sphere.Position(ray)
				if position > 0 && position < minPosition {
					minIndex = i
					minPosition = position
				}
			}

			if minIndex >= 0 {
				sphere := s.spheres[minIndex]
				N := ray.PointAt(minPosition).Sub(sphere.center)
				r := uint8(math.Abs(Normalize(N).InnerProduct(Normalize(sphere.center.Sub(s.camera.pos)))) * 255)
				if minIndex == 1 {
					println(r)
				}
				img.Set(j, s.height-i, color.RGBA{
					r,
					r,
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
