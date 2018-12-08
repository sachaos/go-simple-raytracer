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

type Metal struct {
	albedo *Vector
}

func (m *Metal) Scatter(ray *Ray, h *Hit) (*Ray, *Vector) {
	dir := ray.Dir.Sub(h.N.Multi(ray.Dir.InnerProduct(h.N)).Multi(2))
	return &Ray{Origin: h.P, Dir: Normalize(dir)}, m.albedo
}
