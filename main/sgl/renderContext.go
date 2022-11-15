package sgl

import "math"

type RenderContext struct {
	*Bitmap
	scanBuffer []int
}

func NewRenderContext(width, height int) *RenderContext {
	bitmap := NewBitmap(width, height)
	rc := RenderContext{
		Bitmap:     bitmap,
		scanBuffer: make([]int, height*2),
	}
	return &rc
}

func (r *RenderContext) FillTriangle(v1, v2, v3 Vertex) {
	screenSpaceTransform := NewMat4f().InitScreenSpaceTransform(float32(r.width)/2, float32(r.height)/2)
	minYVert := v1.Transform(screenSpaceTransform).PerspectiveDivide()
	midYVert := v2.Transform(screenSpaceTransform).PerspectiveDivide()
	maxYVert := v3.Transform(screenSpaceTransform).PerspectiveDivide()

	if maxYVert.Pos.Y < midYVert.Pos.Y {
		maxYVert, midYVert = midYVert, maxYVert
	}
	if midYVert.Pos.Y < minYVert.Pos.Y {
		midYVert, minYVert = minYVert, midYVert
	}
	if maxYVert.Pos.Y < midYVert.Pos.Y {
		maxYVert, midYVert = midYVert, maxYVert
	}

	r.ScanTriangle(minYVert, midYVert, maxYVert, minYVert.CrossProduct(maxYVert, midYVert) >= 0)
}

func (r *RenderContext) ScanTriangle(minYVert, midYVert, maxYVert Vertex, handedness bool) {
	topToBottom := NewEdge(minYVert, maxYVert)
	topToMid := NewEdge(minYVert, midYVert)
	midToBottom := NewEdge(midYVert, maxYVert)

	left := topToBottom
	right := topToMid

	if handedness {
		left, right = right, left
	}

	yStart := topToMid.YStart
	yEnd := topToMid.YEnd

	for j := yStart; j < yEnd; j++ {
		r.DrawScanLine(left, right, j)
		left.Step()
		right.Step()
	}

	left = topToBottom
	right = midToBottom

	if handedness {
		left, right = right, left
	}

	yStart = midToBottom.YStart
	yEnd = midToBottom.YEnd

	for j := yStart; j < yEnd; j++ {
		r.DrawScanLine(left, right, j)
		left.Step()
		right.Step()
	}
}

func (r *RenderContext) DrawScanLine(left, right *Edge, j int) {
	xMin := int(math.Ceil(float64(left.X)))
	xMax := int(math.Ceil(float64(right.X)))

	for i := xMin; i < xMax; i++ {
		r.DrawPixel(i, j, 200, 200, 200, 255)
	}
}
