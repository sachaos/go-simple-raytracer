package main

func main() {
	height := 200
	width := 200

	lookFrom := NewVector(-2, 0, 0)
	lookAt := NewVector(2, 0, 0)
	vup := NewVector(0, 0, 1)
	c := NewCamera(lookFrom, lookAt, vup, 0.5, 1.0)

	sphere := &Sphere{NewVector(0, 0.0, 0.2), 0.2}
	scene := &Scene{height: height, width: width, camera: c, sphere: sphere}

	scene.Render()
}
