package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
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
	sampleCount   int
}

func (s Scene) background(ray *Ray) *Vector {
	return &Vector{x: 255.0, y: 255.0, z: 255.0}
}

func (s *Scene) Render() {
	img := image.NewRGBA(image.Rect(0, 0, s.width, s.height))

	for i := 0; i < s.height; i++ {
		for j := 0; j < s.width; j++ {
			c := &Vector{}
			for k := 0; k < s.sampleCount; k++ {
				u := (float64(j)+rand.Float64())/float64(s.width)
				v := (float64(i)+rand.Float64())/float64(s.height)
				ray := s.camera.GetRay(u, v)

				hit := s.objectList.Position(ray)
				if hit != nil {
					red := math.Abs(Normalize(hit.N).InnerProduct(Normalize(hit.P.Sub(s.camera.pos)))) * 255
					c = c.Add(&Vector{x: red})
				} else {
					c = c.Add(s.background(ray))
				}
			}
			println(c.String())
			c = c.Multi(1.0 / float64(s.sampleCount))

			img.Set(j, s.height-i, color.RGBA{
				uint8(c.x),
				uint8(c.y),
				uint8(c.z),
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
