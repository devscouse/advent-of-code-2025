package common

import "math"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{X: x, Y: y, Z: z}
}

func (v *Vector3) Sub(other *Vector3) *Vector3 {
	return NewVector3(
		v.X-other.X,
		v.Y-other.Y,
		v.Z-other.Z,
	)
}

func (v *Vector3) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v *Vector3) EuclideanDistance(other *Vector3) float64 {
	delta := other.Sub(v)
	return delta.Magnitude()
}
