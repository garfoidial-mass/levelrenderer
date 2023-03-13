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

type Texture struct {
	id      uint32
	surface *sdl.Surface
}

// checks if the error is nil or not; panics the program if err exists, else does nothing
func testErr(err error) {
	if err != nil {
		panic(err)
	}
}

var textures map[string]*Texture
var renderer *sdl.Renderer

func loadTexture(filepath string, texname string) { // partially from http://www.sdltutorials.com/sdl-tip-sdl-surface-to-opengl-texture
	newsurf, err := img.Load(filepath)
	testErr(err)
	texture := new(Texture)
	finalsurf, err := newsurf.ConvertFormat(uint32(sdl.PIXELFORMAT_RGBA32), 0)
	testErr(err)
	texture.surface = finalsurf
	newsurf.Free()
	gl.GenTextures(1, &texture.id)
	gl.BindTexture(gl.TEXTURE_2D, texture.id)

	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, texture.surface.W, texture.surface.H, 0, gl.RGBA, gl.UNSIGNED_BYTE, texture.surface.Data())

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	textures[texname] = texture
}

var time uint64
var deltatime uint64

func main() {
	fmt.Println("help")
	// init sdl
	err := sdl.Init(sdl.INIT_EVERYTHING)
	testErr(err)
	// create the window to draw on
	window, err := sdl.CreateWindow("Test window", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 800, 600, sdl.WINDOW_OPENGL)
	testErr(err)
	// creates an opengl context to allow drawing to the window
	context, err := window.GLCreateContext()
	testErr(err)
	fmt.Println("context")
	renderer, err = sdl.CreateRenderer(window, -1, 0)
	testErr(err)
	if renderer != nil {
		fmt.Println("renderer")
	}
	fmt.Println("renderer created")

	textures = make(map[string]*Texture, 0)

	//initialize opengl library
	err = gl.Init()
	testErr(err)

	fmt.Println("opengl")

	// perform basic opengl set up to begin drawing
	gl.Enable(gl.DEPTH_TEST)
	gl.Enable(gl.TEXTURE_2D)
	gl.ClearColor(0, 0, 0, 0)
	gl.DepthFunc(gl.LEQUAL)
	gl.Viewport(0, 0, 800, 600)

	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	gl.Frustum(-0.1, 0.1, -0.1, 0.1, 0.1, 100)
	gl.MatrixMode(gl.MODELVIEW)

	gl.ActiveTexture(gl.TEXTURE0)

	loadTexture("templesky.gif", "ALENA")
	gl.BindTexture(gl.TEXTURE_2D, textures["ALENA"].id)
	testErr(err)
	lasttime := uint64(0)
	time := sdl.GetTicks64()
	running := true
	for running {
		time = sdl.GetTicks64()
		deltatime = time - lasttime
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
		lasttime = time

	}
	sdl.GLDeleteContext(context)
	window.Destroy()
	sdl.Quit()
}

var angle float32 = 0

func draw() {
	gl.Clear(gl.DEPTH_BUFFER_BIT | gl.COLOR_BUFFER_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.LoadIdentity()
	gl.Translatef(0, 0, -4)
	gl.Rotatef(angle, 0, 1, 0)
	angle += 0.1
	gl.Begin(gl.TRIANGLES)
	gl.TexCoord2f(0, 1)
	//gl.Color3f(0.1, 1.0, 0.5)
	gl.Vertex2f(-1.0, -1.0)
	gl.TexCoord2f(0.5, 0)
	//gl.Color3f(0.1, 1.0, 0.5)
	gl.Vertex2f(0.0, 1.0)
	gl.TexCoord2f(1, 1)
	//gl.Color3f(0.1, 1.0, 0.5)
	gl.Vertex2f(1, -1.0)

	gl.End()
}
