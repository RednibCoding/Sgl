package sgl

type Bitmap struct {
	width      int
	height     int
	components []byte
}

func NewBitmap(width, height int) *Bitmap {
	return &Bitmap{
		width:      width,
		height:     height,
		components: make([]byte, width*height*4),
	}
}

func (b *Bitmap) Clear(shade byte) {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			index := (x + y*b.width) * 4
			// Alpha component is always max
			b.components[index+0] = shade // r
			b.components[index+1] = shade // g
			b.components[index+2] = shade // b
			b.components[index+3] = 255   // a
		}
	}
}

func (bi *Bitmap) DrawPixel(x, y int, r, g, b, a byte) {
	index := (y*bi.width + x) * 4
	if index < len(bi.components)-4 && index >= 0 {
		bi.components[index+0] = r
		bi.components[index+1] = g
		bi.components[index+2] = b
		bi.components[index+3] = a
	}
}

func (b *Bitmap) Width() int {
	return b.width
}

func (b *Bitmap) Height() int {
	return b.height
}
