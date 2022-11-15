package sgl

import "math"

type Edge struct {
	X      float32
	xStep  float32
	YStart int
	YEnd   int
}

func NewEdge(minYVert, maxYVert Vertex) *Edge {
	yDist := maxYVert.Pos.Y - minYVert.Pos.Y
	xDist := maxYVert.Pos.X - minYVert.Pos.X

	yStart := int(math.Ceil(float64(minYVert.Pos.Y)))
	yEnd := int(math.Ceil(float64(maxYVert.Pos.Y)))
	xStep_ := float32(xDist) / float32(yDist)
	yPrestep := float32(yStart) - minYVert.Pos.Y
	x := minYVert.Pos.X + yPrestep*xStep_

	e := Edge{
		X:      x,
		xStep:  xStep_,
		YStart: yStart,
		YEnd:   yEnd,
	}
	return &e
}

func (e *Edge) Step() {
	e.X += e.xStep
}
