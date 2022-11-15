package sgl

type Vertex struct {
	Pos   Vec4f
	Color Vec4f
}

func NewVertex(x, y, z float32) Vertex {
	v := Vertex{Pos: NewVec4f(x, y, z, 1)}
	return v
}

func (v Vertex) Transform(transformMat Mat4f) Vertex {
	pos := transformMat.Transform(v.Pos)
	return Vertex{Pos: pos}
}

func (v Vertex) PerspectiveDivide() Vertex {
	return Vertex{Pos: NewVec4f(v.Pos.X/v.Pos.W, v.Pos.Y/v.Pos.W, v.Pos.Z/v.Pos.W, v.Pos.W)}
}

func (v Vertex) CrossProduct(b, c Vertex) float32 {
	x1 := b.Pos.X - v.Pos.X
	y1 := b.Pos.Y - v.Pos.Y

	x2 := c.Pos.X - v.Pos.X
	y2 := c.Pos.Y - v.Pos.Y

	return (x1*y2 - x2*y1)
}
