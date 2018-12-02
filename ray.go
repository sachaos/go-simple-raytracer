package main

import (
	"gonum.org/v1/gonum/mat"
)

type Ray struct {
	Origin *mat.Dense
	Dir    *mat.Dense
}

func (r *Ray) PointAt(t float64) *mat.Dense {
	scaledDir := mat.NewDense(3, 1, make([]float64, 3))
	scaledDir.Scale(t, r.Dir)

	res := mat.NewDense(3, 1, make([]float64, 3))
	res.Add(r.Origin, scaledDir)

	return res
}
