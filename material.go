package main

type Material interface {
	Scatter(ray *Ray, h *Hit) (*Ray, *Vector)
}

type Lambertian struct {
	albedo *Vector
}

func (l *Lambertian) Scatter(ray *Ray, h *Hit) (*Ray, *Vector) {
	target := h.P.Add(h.N).Add(NewRandomVectorInUnitSphere())
	return &Ray{Origin: h.P, Dir: target.Sub(h.P)}, l.albedo
}
