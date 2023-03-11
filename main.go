package main

import (
	"fmt"
	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/sdl"
)

// checks if the error is nil or not; panics the program if err exists, else does nothing
func testErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	// init sdl
	err := sdl.Init(sdl.INIT_EVERYTHING)
	testErr(err)
	// create the window to draw on
	window, err := sdl.CreateWindow("Test window", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, sdl.WINDOW_OPENGL)
	testErr(err)
	// creates an opengl context to allow drawing to the window
	context, err := window.GLCreateContext()
	testErr(err)

	//initialize opengl library
	err = gl.Init()
	testErr(err)

	// perform basic opengl set up to begin drawing
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0,0,0,0)
	gl.DepthFunc(gl.LEQUAL)

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				fmt.Println("quitevent")
				break
			}

		}

	}
	sdl.GLDeleteContext(context)
	window.Destroy()
	sdl.Quit()
}
