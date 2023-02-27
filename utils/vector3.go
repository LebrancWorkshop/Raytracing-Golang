package utils

import "math"

type Vector3 struct {
	X float64
	Y float64
	Z float64
}

func (v Vector3) AddVect(vd Vector3) Vector3 {
	return Vector3{v.X + vd.X, v.Y + vd.Y, v.Z + vd.Z}
}

func (v Vector3) SubVect(vd Vector3) Vector3 {
	return Vector3{v.X - vd.X, v.Y - vd.Y, v.Z - vd.Z}
}

func (v Vector3) MulScal(scale float64) Vector3 {
	return Vector3{v.X * scale, v.Y * scale, v.Z * scale}
}

func (v Vector3) DivScal(scale float64) Vector3 {
	return Vector3{v.X / scale, v.Y / scale, v.Z / scale}
}

func (v Vector3) LenSqVect() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector3) LenVect() float64 {
	return math.Sqrt(v.LenSqVect())
}

func (v Vector3) DotVect(v1 Vector3, v2 Vector3) float64 {
	return (v1.X * v2.X) + (v1.Y * v2.Y) + (v1.Z * v2.Z)
}

func (v Vector3) UnitVect() Vector3 {
	return Vector3{v.X / v.LenVect(), v.Y / v.LenVect(), v.Z / v.LenVect()}
}