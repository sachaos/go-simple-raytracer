package main

import "fmt"

type Ray struct {
	Origin *Vector
	Dir    *Vector
}

func (r *Ray) PointAt(t float64) *Vector {
	return r.Dir.Multi(t).Add(r.Origin)
}

func (r *Ray) String() string {
	return fmt.Sprintf("Origin: %v, Dir: %v", r.Origin, r.Dir)
}
