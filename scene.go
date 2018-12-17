package main

import (
	"fmt"
	"github.com/mono0x/prand"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"sync"
	"time"
)

type Hit struct {
	T        float64
	P        *Vector
	N        *Vector
	Material Material
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
	N := Normalize(P.Sub(sphere.center))

	return &Hit{
		T:        minPosition,
		P:        P,
		N:        N,
		Material: sphere.Material,
	}
}

type Scene struct {
	height, width int
	camera        *Camera
	objectList    *ObjectList
	sampleCount   int
}

func (s *Scene) background(ray *Ray) *Vector {
	return &Vector{x: 255.0, y: 255.0, z: 255.0}
}

func (s *Scene) color(ray *Ray) *Vector {
	hit := s.objectList.Position(ray)
	if hit != nil {
		nextRay, albedo := hit.Material.Scatter(ray, hit)
		return s.color(nextRay).MultiPerElem(albedo)
	} else {
		return s.background(ray)
	}
}

func (s *Scene) Render() {
	img := image.NewRGBA(image.Rect(0, 0, s.width, s.height))

	maxGoroutineNum := 8

	var wg sync.WaitGroup
	for goroutineNum := 0; goroutineNum < maxGoroutineNum; goroutineNum++ {
		heightStart := goroutineNum * (s.height / maxGoroutineNum)
		heightEnd := (goroutineNum + 1) * (s.height / maxGoroutineNum)
		wg.Add(1)
		go func(heightStart, heightEnd, num int) {
			start := time.Now()
			fmt.Printf("num %d start at: %s\n", num, start)
			defer wg.Done()
			defer func() { fmt.Printf("num %d end: %s\n", num, time.Now().Sub(start)) }()
			for i := heightStart; i < heightEnd; i++ {
				for j := 0; j < s.width; j++ {
					c := &Vector{}
					for k := 0; k < s.sampleCount; k++ {
						u := (float64(j) + prand.Float64()) / float64(s.width)
						v := (float64(i) + prand.Float64()) / float64(s.height)
						ray := s.camera.GetRay(u, v)

						c = c.Add(s.color(ray))
					}
					c = c.Multi(1.0 / float64(s.sampleCount))

					img.Set(j, s.height-i, color.RGBA{
						uint8(c.x),
						uint8(c.y),
						uint8(c.z),
						255,
					})
				}
			}
		}(heightStart, heightEnd, goroutineNum)
	}

	wg.Wait()
	f, err := os.OpenFile("out.png", os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
