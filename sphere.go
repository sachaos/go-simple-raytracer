package main

import "math"

type Sphere struct{
	center *Vector
	r float64
	Material Material
}

func (s *Sphere)Position(ray *Ray) float64 {
	a := ray.Dir.InnerProduct(ray.Dir)
	b := 2.0 * ray.Dir.InnerProduct(ray.Origin.Sub(s.center))
	c := ray.Origin.Sub(s.center).InnerProduct(ray.Origin.Sub(s.center)) - s.r * s.r

	D := b*b - 4*a*c
	if (D < 0) {
		return -1
	}

	return (-b - math.Sqrt(D)) / (2 * a)
}

func (s *Sphere) IsConflict(xs *Sphere) bool {
	d := Len(s.center.Sub(xs.center))
	return d < (s.r + xs.r)
}
