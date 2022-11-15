package sgl

import (
	"fmt"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

type Display struct {
	title         string
	window        *sdl.Window
	renderer      *sdl.Renderer
	renderSurface *sdl.Texture
	frameBuffer   *RenderContext
}

func NewDisplay(title string, width int, height int) *Display {
	display := Display{title: title}
	display.init(width, height)
	return &display
}

func (s *Display) init(width, height int) {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}

	window, err := sdl.CreateWindow(s.title, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, int32(width), int32(height), sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.window = window

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.renderer = renderer

	surface, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(width), int32(height))
	if err != nil {
		fmt.Println(err)
		return
	}
	s.renderSurface = surface

	s.frameBuffer = NewRenderContext(width, height)
}

func (s *Display) SwapBuffers() {
	s.renderSurface.Update(nil, unsafe.Pointer(&s.frameBuffer.components[0]), s.frameBuffer.width*4)
	s.renderer.Copy(s.renderSurface, nil, nil)
	s.renderer.Present()
}

func (s *Display) Destroy() {
	if s.renderSurface != nil {
		s.renderSurface.Destroy()
	}
	if s.renderer != nil {
		s.renderer.Destroy()
	}
	if s.window != nil {
		s.window.Destroy()
	}
}

func (s *Display) Clear(shade byte) {
	s.frameBuffer.Clear(shade)
}

func (s *Display) DrawPixel(x, y int, r, g, b, a byte) {
	s.frameBuffer.DrawPixel(x, y, r, g, b, a)
}

func (s *Display) FrameBuffer() *RenderContext {
	return s.frameBuffer
}

func (s *Display) Title() string {
	return s.title
}

func (s *Display) Width() int {
	return s.frameBuffer.width
}

func (s *Display) Height() int {
	return s.frameBuffer.height
}
