package main

func main() {
	height := 400
	width := 400

	lookFrom := NewVector(-2, 0, 0.3)
	lookAt := NewVector(2, 0, 0)
	vup := NewVector(0, 0, 1)
	c := NewCamera(lookFrom, lookAt, vup, 0.5, 1.0)

	sphere := &Sphere{NewVector(0, 0, 0.2), 0.2}
	ground := &Sphere{NewVector(0, 0, -100), 100}

	objectList := &ObjectList{spheres: []*Sphere{sphere, ground}}
	scene := &Scene{sampleCount: 100, height: height, width: width, camera: c, objectList: objectList}

	scene.Render()
}
