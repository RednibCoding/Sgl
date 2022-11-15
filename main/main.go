package main

import (
	"main/sgl"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Sgl"
var winWidth, winHeight int = 1280, 720

func main() {
	display := sgl.NewDisplay(winTitle, winWidth, winHeight)
	defer display.Destroy()
	target := display.FrameBuffer()

	minYVert := sgl.NewVertex(-1, -1, 0)
	midYVert := sgl.NewVertex(0, 1, 0)
	maxYVert := sgl.NewVertex(1, -1, 0)

	aspect := float32(target.Width()) / float32(target.Height())
	projection := sgl.NewMat4f().InitPerspective(70, aspect, 0.1, 1000)
	var rotCounter float32 = 0

	previousTime := time.Now().UnixMilli()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}

		currentTime := time.Now().UnixMilli()
		var deltaTime float32 = float32(currentTime-previousTime) / 1000.0
		previousTime = currentTime

		rotCounter += deltaTime
		translation := sgl.NewMat4f().InitTranslation(0, 0, 3)
		rotation := sgl.NewMat4f().InitRotationF(0, rotCounter, 0)
		transform := projection.Mul(translation.Mul(rotation))

		target.Clear(0)
		target.FillTriangle(maxYVert.Transform(transform), midYVert.Transform(transform), minYVert.Transform(transform))

		display.SwapBuffers()
	}
}
