package main

import (
	"github.com/pkg/profile"
	"math/rand"
	"time"
)

var genRange = 10.0
var sizeMin = 0.05
var sizeMax = 0.3

func genSphere(currentSpheres []*Sphere) *Sphere {
	x := rand.Float64() * genRange * 2 - genRange
	y := rand.Float64() * genRange * 2 - genRange

	size := (sizeMax - sizeMin) * rand.Float64() + sizeMin
	c := &Vector{rand.Float64(), rand.Float64(), rand.Float64()}
	materialType := rand.Intn(2)
	var material Material
	switch materialType {
	case 0:
		material = &Lambertian{albedo: c}
	case 1:
		material = &Metal{albedo: c}
	}
	s := &Sphere{NewVector(x, y, size), size, material}
	for _, xs := range currentSpheres {
		if s.IsConflict(xs) {
			return nil
		}
	}

	return s
}

func main() {
	defer profile.Start(profile.ProfilePath("."), profile.MutexProfile).Stop()
	rand.Seed(time.Now().UnixNano())

	height := 750
	width := 1000

	lookFrom := NewVector(-4.5, 0, 0.8)
	lookAt := NewVector(0, 0, 0.25)
	vup := NewVector(0, 0, 1)
	c := NewCamera(lookFrom, lookAt, vup, 0.5, float64(width) / float64(height))

	var spheres []*Sphere
	for i := 0; i < 300; i++ {
		s := genSphere(spheres)
		if s != nil {
			spheres = append(spheres, s)
		}
	}

	ground := &Sphere{NewVector(0, 0, -10000), 10000, &Lambertian{
		albedo: &Vector{0.7, 0.8, 0.8},
	}}
	spheres = append(spheres, ground)

	objectList := &ObjectList{spheres: spheres}
	scene := &Scene{sampleCount: 100, height: height, width: width, camera: c, objectList: objectList}

	scene.Render()
}
