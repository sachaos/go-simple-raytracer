package main

import (
	"fmt"
	"math"
)

type Vector struct {
	x, y, z float64
}

func NewVector(x, y, z float64) *Vector {
	return &Vector{x, y, z}
}

func (v *Vector) Add(tv *Vector) *Vector {
	return &Vector{v.x + tv.x, v.y + tv.y, v.z + tv.z}
}

func (v *Vector) Sub(tv *Vector) *Vector {
	return &Vector{v.x - tv.x, v.y - tv.y, v.z - tv.z}
}

func (v *Vector) Multi(a float64) *Vector {
	return &Vector{v.x * a, v.y * a, v.z * a}
}

func (v *Vector) InnerProduct(tv *Vector) float64 {
	return v.x*tv.x + v.y*tv.y + v.z*tv.z
}

func (v *Vector) OuterProduct(tv *Vector) *Vector {
	return &Vector{v.y*tv.z - v.z*tv.y, v.z*tv.x - v.x*tv.z, v.x*tv.y - v.y*tv.x}
}

func (v *Vector) String() string {
	return fmt.Sprintf("(%f, %f, %f)", v.x, v.y, v.z)

}

func (v *Vector) Lerp(tv *Vector, alpha float64) *Vector{
	return &Vector{
		v.x + (tv.x-v.x)*alpha,
		v.y + (tv.y-v.y)*alpha,
		v.z + (tv.z-v.z)*alpha,
	}
}

func Len(v *Vector) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
}

func Normalize(v *Vector) *Vector {
	return v.Multi(1.0 / Len(v))
}
