package sgl

import (
	"fmt"
	"math"
)

type Vec4f struct {
	X float32
	Y float32
	Z float32
	W float32
}

func NewVec4f(x, y, z, w float32) Vec4f {
	v := Vec4f{X: x, Y: y, Z: z, W: w}
	return v
}

func (v Vec4f) Len() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z + v.W*v.W)))
}

func (v Vec4f) Max() float32 {
	return float32(math.Max(math.Max(float64(v.X), float64(v.Y)), math.Max(float64(v.Z), float64(v.W))))
}

func (v Vec4f) Dot(r Vec4f) float32 {
	return v.X*r.X + v.Y*r.Y + v.Z*r.Z + v.W*r.W
}

func (v Vec4f) Cross(r Vec4f) Vec4f {
	x := v.Y*r.Z - v.Z*r.Y
	y := v.Z*r.X - v.X*r.Z
	z := v.X*r.Y - v.Y*r.X
	return Vec4f{X: x, Y: y, Z: z, W: 0}
}

func (v Vec4f) Normalized() Vec4f {
	len := v.Len()
	return Vec4f{X: v.X / len, Y: v.Y / len, Z: v.Z / len, W: v.W / len}
}

func (v Vec4f) Rotate(axis Vec4f, angle float32) Vec4f {
	sinAngle := float32(math.Sin(-float64(angle)))
	cosAngle := float32(math.Cos(-float64(angle)))

	return v.Cross(axis.MulF(sinAngle)).AddV((v.MulF(cosAngle)).AddV(axis.MulF(v.Dot(axis.MulF(1 - cosAngle)))))
}

func (v Vec4f) Lerp(dest Vec4f, lerpFactor float32) Vec4f {
	return dest.SubV(v).MulF(lerpFactor).AddV(v)
}

func (v Vec4f) AddV(r Vec4f) Vec4f {
	return Vec4f{X: v.X + r.X, Y: v.Y + r.Y, Z: v.Z + r.Z, W: v.W + r.W}
}

func (v Vec4f) AddF(r float32) Vec4f {
	return Vec4f{X: v.X + r, Y: v.Y + r, Z: v.Z + r, W: v.W + r}
}

func (v Vec4f) SubV(r Vec4f) Vec4f {
	return Vec4f{X: v.X - r.X, Y: v.Y - r.Y, Z: v.Z - r.Z, W: v.W - r.W}
}

func (v Vec4f) SubF(r float32) Vec4f {
	return Vec4f{X: v.X - r, Y: v.Y - r, Z: v.Z - r, W: v.W - r}
}

func (v Vec4f) MulV(r Vec4f) Vec4f {
	return Vec4f{X: v.X * r.X, Y: v.Y * r.Y, Z: v.Z * r.Z, W: v.W * r.W}
}

func (v Vec4f) MulF(r float32) Vec4f {
	return Vec4f{X: v.X * r, Y: v.Y * r, Z: v.Z * r, W: v.W * r}
}

func (v Vec4f) DivV(r Vec4f) Vec4f {
	return Vec4f{X: v.X / r.X, Y: v.Y / r.Y, Z: v.Z / r.Z, W: v.W / r.W}
}

func (v Vec4f) DivF(r float32) Vec4f {
	return Vec4f{X: v.X / r, Y: v.Y / r, Z: v.Z / r, W: v.W / r}
}

func (v Vec4f) Abs() Vec4f {
	return Vec4f{
		X: float32(math.Abs(float64(v.X))),
		Y: float32(math.Abs(float64(v.Y))),
		Z: float32(math.Abs(float64(v.Z))),
		W: float32(math.Abs(float64(v.W))),
	}
}

func (v Vec4f) Equals(r Vec4f) bool {
	return v.X == r.X && v.Y == r.Y && v.Z == r.Z && v.W == r.W
}

func (v Vec4f) ToString() string {
	return "(" + fmt.Sprintf("%f", v.X) + ", " + fmt.Sprintf("%f", v.Y) + ", " + fmt.Sprintf("%f", v.Z) + ", " + fmt.Sprintf("%f", v.W) + ")"
}
