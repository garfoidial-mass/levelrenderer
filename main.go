package main

import (
	"fmt"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// defines a 2d point that the walls go between
type vector2 struct {
	x int
	y int
}

// defines a "line" or vertical wall that goes between 2 points and has a top, middle, and bottom texture
type lineDef struct {
	top    string // top texture
	center string // middle texture
	bottom string // bottom texture

	start vector2 // start point
	end   vector2 // end point
}

// defines a "sector" or collection of lines with a textured floor and ceiling
type sector struct {
	ceilingHeight int
	floorHeight   int

	ceilingTex string
	floorTex   string

	lines []lineDef
}

// checks if the error is nil or not; panics the program if err exists, else does nothing
func testErr(err error) {
	if err != nil {
		panic(err)
	}
}

var textures map[string]*sdl.Texture
var renderer *sdl.Renderer

func loadTexture(filepath string, texname string) {
	newtex, err := img.LoadTexture(renderer, filepath)
	testErr(err)
	textures[texname] = newtex
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
	renderer, err = sdl.CreateRenderer(window, -1, 0)
	testErr(err)
	if renderer != nil {
		print("renderer\n")
	}

	loadTexture("C:\\Users\\Basil\\Pictures\\bigpics\\the photoshoot\\DSC_0010.JPG", "ALENA")

	//initialize opengl library
	err = gl.Init()
	testErr(err)

	// perform basic opengl set up to begin drawing
	gl.Enable(gl.DEPTH_TEST)
	gl.ClearColor(0, 0, 0, 0)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, 800, 600)
	gl.ActiveTexture(gl.TEXTURE0)
	textures["ALENA"].GLBind(nil, nil)

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
		draw()          // render the 3d scene on screen
		window.GLSwap() //rendering is double buffered, so you have to swap them to get the new image displayed

	}
	sdl.GLDeleteContext(context)
	window.Destroy()
	sdl.Quit()
}

func draw() {
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	gl.Begin(gl.TRIANGLES)
	gl.Color3f(0.1, 1.0, 0.5)
	gl.Vertex2f(-1.0, -1.0)
	gl.Color3f(0.1, 1.0, 0.5)
	gl.Vertex2f(0.0, 1.0)
	gl.Color3f(0.1, 1.0, 0.5)
	gl.Vertex2f(1, -1.0)

	gl.End()
}
