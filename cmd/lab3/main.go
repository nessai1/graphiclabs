package main

import (
	"fmt"
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
	app := application.CreateApplication("Lab 3", appWight, appHeight)
	app.Program = initOpenGL()
	app.Handler = programHandler
	app.Mode = gl.LINE_STRIP
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

	spline := figures.Spline{}
	for {
		if window.GetKey(glfw.KeyEscape) == 1 {
			shouldClose <- true
		}

		if window.GetMouseButton(glfw.MouseButton1) == 1 {
			x, y := window.GetCursorPos()
			fmt.Println("Point: ", figures.GeneratePointByWindow(appHeight, appWight, x, y))

			spline.AddPoint(figures.GeneratePointByWindow(appHeight, appWight, x, y))
			fmt.Println(spline.Counter)
			if spline.Counter > 1 {
				var f figures.Figure = &spline

				resultFigures := []*figures.Figure{&f}
				figuresChannel <- resultFigures
			}

			time.Sleep(time.Millisecond * 400)
		}

		//time.Sleep(time.Millisecond * 5)
	}

}
