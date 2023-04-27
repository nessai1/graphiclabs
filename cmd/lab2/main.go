package main

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"graphiclabs/internal/application"
	"graphiclabs/internal/figures"
	"log"
	"time"
)

const (
	vertexShaderSource = `
    #version 410
    in vec3 vp;
    void main() {
        gl_Position = vec4(vp, 1.0);
    }
` + "\x00"

	fragmentShaderSource = `
    #version 410
    out vec4 frag_colour;
    void main() {
        frag_colour = vec4(0.1, 0.97, 0.03, 1);
    }
` + "\x00"
	appHeight = 600
	appWight  = 600
)

func main() {
	app := application.CreateApplication("Lab 1", appWight, appHeight)
	app.Program = initOpenGL()
	app.Handler = programHandler
	circle := &figures.CircleFragment{
		Point:    figures.Point{X: 0, Y: 0},
		Rotation: 0,
		Radius:   0.3,
	}
	var f figures.Figure = circle
	app.Figures = []*figures.Figure{&f}
	app.Run()
}

// initOpenGL initializes OpenGL and returns an intiialized program.
func initOpenGL() uint32 {
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	vertexShader, err := application.CompileShader(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}
	fragmentShader, err := application.CompileShader(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vertexShader)
	gl.AttachShader(prog, fragmentShader)
	gl.LinkProgram(prog)
	return prog
}

func programHandler(window *glfw.Window, shouldClose chan bool, figuresChannel chan []*figures.Figure) {
	rotation := 0
	for {
		if window.GetKey(glfw.KeyEscape) == 1 {
			shouldClose <- true
		}

		if window.GetKey(glfw.KeyLeft) == 1 {
			rotation -= 1
			circle := &figures.CircleFragment{
				Point:    figures.Point{X: 0, Y: 0},
				Rotation: rotation,
				Radius:   0.3,
			}

			var f figures.Figure = circle

			resultFigures := []*figures.Figure{&f}
			figuresChannel <- resultFigures
		}

		if window.GetKey(glfw.KeyRight) == 1 {
			rotation += 1
			circle := &figures.CircleFragment{
				Point:    figures.Point{X: 0, Y: 0},
				Rotation: rotation,
				Radius:   0.3,
			}

			var f figures.Figure = circle

			resultFigures := []*figures.Figure{&f}
			figuresChannel <- resultFigures
		}

		if window.GetKey(glfw.KeyLeft) == 1 {
			log.Println("LEFT KEY")
		}

		time.Sleep(time.Millisecond * 5)
	}

}
