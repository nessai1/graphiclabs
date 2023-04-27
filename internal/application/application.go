package application

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"graphiclabs/internal/figures"
	"graphiclabs/internal/window"
	"runtime"
	"strings"
)

const (
	width  = 800
	height = 600
)

type Application struct {
	shouldClose    chan bool
	figuresChannel chan []*figures.Figure

	Window   *glfw.Window
	Program  uint32
	vao      uint32
	Figures  []*figures.Figure
	Handler  func(*glfw.Window, chan bool, chan []*figures.Figure)
	vertices []float32
	Mode     uint32
}

func CreateApplication(name string, inputWidth int, inputHeight int) *Application {
	if inputWidth == 0 {
		inputWidth = width
	}

	if inputHeight == 0 {
		inputHeight = height
	}

	application := new(Application)
	if createdWindow, err := window.InitWindow(inputWidth, inputHeight, name); err != nil {
		panic(err)
	} else {
		application.Window = createdWindow
	}

	application.Mode = gl.TRIANGLES

	return application
}

func (app *Application) Run() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	app.compileVertices()
	if len(app.vertices) > 0 {
		app.vao = app.makeVao()
	}
	app.shouldClose = make(chan bool)
	app.figuresChannel = make(chan []*figures.Figure)

	go app.Handler(app.Window, app.shouldClose, app.figuresChannel)

	for !app.Window.ShouldClose() {
		select {
		case <-app.shouldClose:
			app.Window.SetShouldClose(true)
		case app.Figures = <-app.figuresChannel:
			app.compileVertices()
			app.vao = app.makeVao()
			app.draw()
		default:
			app.draw()
		}
	}

}

func (app *Application) draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(app.Program)
	gl.BindVertexArray(app.vao)
	gl.DrawArrays(app.Mode, 0, int32(len(app.vertices)))

	glfw.PollEvents()
	app.Window.SwapBuffers()
}

func (app *Application) makeVao() uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(app.vertices), gl.Ptr(app.vertices), gl.STATIC_DRAW) // 4 bytes from float32

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
}

func (app *Application) compileVertices() {
	app.vertices = []float32{}
	for _, figure := range app.Figures {
		var vertices *[]float32
		vertices = (*figure).GetVertices()

		app.vertices = append(app.vertices, *vertices...)
	}
}

func CompileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}
