package main

import (
	"fmt"
	"math"
)

type Camera struct {
	pos, u, v, w, lowerLeft, horizontal, vertical *Vector
}

func NewCamera(lookfrom, lookat, vup *Vector, vfov, aspect float64) *Camera {
	halfHeight := math.Tan(vfov / 2)
	halfWidth := aspect * halfHeight

	var pos *Vector
	pos = lookfrom

	w := Normalize(lookfrom.Sub(lookat))
	u := Normalize(vup.OuterProduct(w))
	v := Normalize(w.OuterProduct(u))

	lowerLeft := lookfrom.Sub(u.Multi(float64(halfWidth))).Sub(v.Multi(float64(halfHeight))).Sub(w)
	horizontal := u.Multi(2.0 * halfWidth)
	vertical := v.Multi(2.0 * halfHeight)

	return &Camera{
		pos, u, v, w,
		lowerLeft, horizontal, vertical,
	}
}

func (c *Camera) GetRay(x, y float64) *Ray {
	return &Ray{Origin: c.pos, Dir: c.lowerLeft.Add(c.horizontal.Multi(x)).Add(c.vertical.Multi(y)).Sub(c.pos)}
}

func (c *Camera) String() string {
	return fmt.Sprintf("pos: %+v, u: %+v, v: %+v, w: %+v\nlowerLeft: %+v, horizontal: %+v, vertical: %+v\n", c.pos, c.u, c.v, c.w, c.lowerLeft, c.horizontal, c.vertical)
}
