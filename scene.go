package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type Hit struct {
	T float64
	P *Vector
	N *Vector
}

type ObjectList struct {
	spheres []*Sphere
}

func (ol *ObjectList) Position(ray *Ray) *Hit {
	minIndex := -1
	minPosition := math.MaxFloat64
	for i, sphere := range ol.spheres {
		position := sphere.Position(ray)
		if position > 0 && position < minPosition {
			minIndex = i
			minPosition = position
		}
	}

	if minIndex < 0 {
		return nil
	}

	sphere := ol.spheres[minIndex]
	P := ray.PointAt(minPosition)
	N := P.Sub(sphere.center)

	return &Hit{
		T: minPosition,
		P: P,
		N: N,
	}
}

type Scene struct {
	height, width int
	camera        *Camera
	objectList    *ObjectList
}

func (s *Scene) Render() {
	img := image.NewRGBA(image.Rect(0, 0, s.width, s.height))

	for i := 0; i < s.height; i++ {
		for j := 0; j < s.width; j++ {
			ray := s.camera.GetRay(float64(j)/float64(s.width), float64(i)/float64(s.height))

			hit := s.objectList.Position(ray)
			if hit == nil {
				continue
			}

			r := uint8(math.Abs(Normalize(hit.N).InnerProduct(Normalize(hit.P.Sub(s.camera.pos)))) * 255)
			img.Set(j, s.height-i, color.RGBA{
				r,
				0,
				0,
				255,
			})
		}
	}
	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
