package main

type Sphere struct{
	center *Vector
	r float64
}

func (s *Sphere)Hit(ray *Ray) bool {
	a := ray.Dir.InnerProduct(ray.Dir)
	b := 2.0 * ray.Dir.InnerProduct(ray.Origin.Sub(s.center))
	c := ray.Origin.Sub(s.center).InnerProduct(ray.Origin.Sub(s.center)) - s.r * s.r

	D := b*b - 4*a*c
	return D > 0
}
