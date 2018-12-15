package main

import "github.com/pkg/profile"

func main() {
	defer profile.Start().Stop()
	height := 750
	width := 1000

	lookFrom := NewVector(-2, 0, 0.3)
	lookAt := NewVector(2, 0, 0)
	vup := NewVector(0, 0, 1)
	c := NewCamera(lookFrom, lookAt, vup, 0.5, 1.5)

	sphere := &Sphere{NewVector(0, -0.25, 0.3), 0.3, &Metal{
		albedo: &Vector{0.95, 0.95, 0.95},
	}}
	sphere2 := &Sphere{NewVector(0, 0.25, 0.2), 0.2, &Lambertian{
		albedo: &Vector{0.5, 0.5, 0.9},
	}}
	sphere3 := &Sphere{NewVector(-0.5, 0, 0.15), 0.15, &Lambertian{
		albedo: &Vector{0.5, 0.9, 0.5},
	}}
	ground := &Sphere{NewVector(0, 0, -100), 100, &Lambertian{
		albedo: &Vector{0.5, 0.5, 0.5},
	}}

	objectList := &ObjectList{spheres: []*Sphere{sphere, sphere2, sphere3, ground}}
	scene := &Scene{sampleCount: 100, height: height, width: width, camera: c, objectList: objectList}

	scene.Render()
}
