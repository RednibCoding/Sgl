package sgl

import "math"

type Mat4f struct {
	m [][]float32
}

func NewMat4f() Mat4f {
	mat := Mat4f{
		m: [][]float32{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
	}
	return mat
}

func (mat Mat4f) InitIdentity() Mat4f {
	mat.m[0][0] = 1
	mat.m[0][1] = 0
	mat.m[0][2] = 0
	mat.m[0][3] = 0
	mat.m[1][0] = 0
	mat.m[1][1] = 1
	mat.m[1][2] = 0
	mat.m[1][3] = 0
	mat.m[2][0] = 0
	mat.m[2][1] = 0
	mat.m[2][2] = 1
	mat.m[2][3] = 0
	mat.m[3][0] = 0
	mat.m[3][1] = 0
	mat.m[3][2] = 0
	mat.m[3][3] = 1
	return mat
}

func (mat Mat4f) InitScreenSpaceTransform(halfW, halfH float32) Mat4f {
	mat.m[0][0] = halfW
	mat.m[0][1] = 0
	mat.m[0][2] = 0
	mat.m[0][3] = halfW
	mat.m[1][0] = 0
	mat.m[1][1] = -halfH
	mat.m[1][2] = 0
	mat.m[1][3] = halfH
	mat.m[2][0] = 0
	mat.m[2][1] = 0
	mat.m[2][2] = 1
	mat.m[2][3] = 0
	mat.m[3][0] = 0
	mat.m[3][1] = 0
	mat.m[3][2] = 0
	mat.m[3][3] = 1

	return mat
}

func (mat Mat4f) InitTranslation(x, y, z float32) Mat4f {
	mat.m[0][0] = 1
	mat.m[0][1] = 0
	mat.m[0][2] = 0
	mat.m[0][3] = x
	mat.m[1][0] = 0
	mat.m[1][1] = 1
	mat.m[1][2] = 0
	mat.m[1][3] = y
	mat.m[2][0] = 0
	mat.m[2][1] = 0
	mat.m[2][2] = 1
	mat.m[2][3] = z
	mat.m[3][0] = 0
	mat.m[3][1] = 0
	mat.m[3][2] = 0
	mat.m[3][3] = 1

	return mat
}

func (mat Mat4f) InitRotationA(x, y, z, angle float32) Mat4f {
	sin := float32(math.Sin(float64(angle)))
	cos := float32(math.Cos(float64(angle)))

	mat.m[0][0] = cos + x*x*(1-cos)
	mat.m[0][1] = x*y*(1-cos) - z*sin
	mat.m[0][2] = x*z*(1-cos) + y*sin
	mat.m[0][3] = 0
	mat.m[1][0] = y*x*(1-cos) + z*sin
	mat.m[1][1] = cos + y*y*(1-cos)
	mat.m[1][2] = y*z*(1-cos) - x*sin
	mat.m[1][3] = 0
	mat.m[2][0] = z*x*(1-cos) - y*sin
	mat.m[2][1] = z*y*(1-cos) + x*sin
	mat.m[2][2] = cos + z*z*(1-cos)
	mat.m[2][3] = 0
	mat.m[3][0] = 0
	mat.m[3][1] = 0
	mat.m[3][2] = 0
	mat.m[3][3] = 1

	return mat
}

func (mat Mat4f) InitRotationF(x, y, z float32) Mat4f {
	rx := NewMat4f()
	ry := NewMat4f()
	rz := NewMat4f()

	rz.m[0][0] = float32(math.Cos(float64(z)))
	rz.m[0][1] = -float32(math.Sin(float64(z)))
	rz.m[0][2] = 0
	rz.m[0][3] = 0
	rz.m[1][0] = float32(math.Sin(float64(z)))
	rz.m[1][1] = float32(math.Cos(float64(z)))
	rz.m[1][2] = 0
	rz.m[1][3] = 0
	rz.m[2][0] = 0
	rz.m[2][1] = 0
	rz.m[2][2] = 1
	rz.m[2][3] = 0
	rz.m[3][0] = 0
	rz.m[3][1] = 0
	rz.m[3][2] = 0
	rz.m[3][3] = 1

	rx.m[0][0] = 1
	rx.m[0][1] = 0
	rx.m[0][2] = 0
	rx.m[0][3] = 0
	rx.m[1][0] = 0
	rx.m[1][1] = float32(math.Cos(float64(x)))
	rx.m[1][2] = -float32(math.Sin(float64(x)))
	rx.m[1][3] = 0
	rx.m[2][0] = 0
	rx.m[2][1] = float32(math.Sin(float64(x)))
	rx.m[2][2] = float32(math.Cos(float64(x)))
	rx.m[2][3] = 0
	rx.m[3][0] = 0
	rx.m[3][1] = 0
	rx.m[3][2] = 0
	rx.m[3][3] = 1

	ry.m[0][0] = float32(math.Cos(float64(y)))
	ry.m[0][1] = 0
	ry.m[0][2] = -float32(math.Sin(float64(y)))
	ry.m[0][3] = 0
	ry.m[1][0] = 0
	ry.m[1][1] = 1
	ry.m[1][2] = 0
	ry.m[1][3] = 0
	ry.m[2][0] = float32(math.Sin(float64(y)))
	ry.m[2][1] = 0
	ry.m[2][2] = float32(math.Cos(float64(y)))
	ry.m[2][3] = 0
	ry.m[3][0] = 0
	ry.m[3][1] = 0
	ry.m[3][2] = 0
	ry.m[3][3] = 1

	mat.m = rz.Mul(ry.Mul(rx)).GetM()

	return mat
}

func (mat Mat4f) InitRotation2V(forward, up Vec4f) Mat4f {
	f := forward.Normalized()
	r := up.Normalized()
	r = r.Cross(f)
	u := f.Cross(r)
	return mat.InitRotation3V(f, u, r)
}

func (mat Mat4f) InitRotation3V(forward, up, right Vec4f) Mat4f {
	f := forward
	r := right
	u := up

	mat.m[0][0] = r.X
	mat.m[0][1] = r.Y
	mat.m[0][2] = r.Z
	mat.m[0][3] = 0
	mat.m[1][0] = u.X
	mat.m[1][1] = u.Y
	mat.m[1][2] = u.Z
	mat.m[1][3] = 0
	mat.m[2][0] = f.X
	mat.m[2][1] = f.Y
	mat.m[2][2] = f.Z
	mat.m[2][3] = 0
	mat.m[3][0] = 0
	mat.m[3][1] = 0
	mat.m[3][2] = 0
	mat.m[3][3] = 1

	return mat
}

func (mat Mat4f) InitScale(x, y, z float32) Mat4f {
	mat.m[0][0] = x
	mat.m[0][1] = 0
	mat.m[0][2] = 0
	mat.m[0][3] = 0
	mat.m[1][0] = 0
	mat.m[1][1] = y
	mat.m[1][2] = 0
	mat.m[1][3] = 0
	mat.m[2][0] = 0
	mat.m[2][1] = 0
	mat.m[2][2] = z
	mat.m[2][3] = 0
	mat.m[3][0] = 0
	mat.m[3][1] = 0
	mat.m[3][2] = 0
	mat.m[3][3] = 1

	return mat
}

func (mat Mat4f) InitPerspective(fov, aspectRatio, zNear, zFar float32) Mat4f {
	rad := (fov * math.Pi) / 180
	tanHalfFOV := float32(math.Tan(float64(rad / 2)))
	zRange := zNear - zFar

	mat.m[0][0] = 1.0 / (tanHalfFOV * aspectRatio)
	mat.m[0][1] = 0
	mat.m[0][2] = 0
	mat.m[0][3] = 0
	mat.m[1][0] = 0
	mat.m[1][1] = 1.0 / tanHalfFOV
	mat.m[1][2] = 0
	mat.m[1][3] = 0
	mat.m[2][0] = 0
	mat.m[2][1] = 0
	mat.m[2][2] = (-zNear - zFar) / zRange
	mat.m[2][3] = 2 * zFar * zNear / zRange
	mat.m[3][0] = 0
	mat.m[3][1] = 0
	mat.m[3][2] = 1
	mat.m[3][3] = 0

	return mat
}

func (mat Mat4f) InitOrthographic(left, right, bottom, top, near, far float32) Mat4f {
	width := right - left
	height := top - bottom
	depth := far - near

	mat.m[0][0] = 2 / width
	mat.m[0][1] = 0
	mat.m[0][2] = 0
	mat.m[0][3] = -(right + left) / width
	mat.m[1][0] = 0
	mat.m[1][1] = 2 / height
	mat.m[1][2] = 0
	mat.m[1][3] = -(top + bottom) / height
	mat.m[2][0] = 0
	mat.m[2][1] = 0
	mat.m[2][2] = -2 / depth
	mat.m[2][3] = -(far + near) / depth
	mat.m[3][0] = 0
	mat.m[3][1] = 0
	mat.m[3][2] = 0
	mat.m[3][3] = 1

	return mat
}

func (mat Mat4f) Transform(r Vec4f) Vec4f {
	return Vec4f{
		X: mat.m[0][0]*r.X + mat.m[0][1]*r.Y + mat.m[0][2]*r.Z + mat.m[0][3]*r.W,
		Y: mat.m[1][0]*r.X + mat.m[1][1]*r.Y + mat.m[1][2]*r.Z + mat.m[1][3]*r.W,
		Z: mat.m[2][0]*r.X + mat.m[2][1]*r.Y + mat.m[2][2]*r.Z + mat.m[2][3]*r.W,
		W: mat.m[3][0]*r.X + mat.m[3][1]*r.Y + mat.m[3][2]*r.Z + mat.m[3][3]*r.W,
	}
}

func (mat Mat4f) Mul(r Mat4f) Mat4f {
	result := NewMat4f()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result.Set(i, j, mat.m[i][0]*r.Get(0, j)+
				mat.m[i][1]*r.Get(1, j)+
				mat.m[i][2]*r.Get(2, j)+
				mat.m[i][3]*r.Get(3, j))
		}
	}
	return result
}

func (mat Mat4f) GetM() [][]float32 {
	result := [][]float32{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			result[i][j] = mat.m[i][j]
		}
	}
	return result
}

func (mat Mat4f) Get(x, y int) float32 {
	return mat.m[x][y]
}

func (mat *Mat4f) SetM(m *[][]float32) {
	mat.m = *m
}

func (mat Mat4f) Set(x, y int, value float32) {
	mat.m[x][y] = value
}
