package application

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"graphiclabs/internal/window"
	"runtime"
	"strings"
)

const (
	width  = 800
	height = 600
)

type Application struct {
	window  *glfw.Window
	Program uint32
	VAO     uint32
	Figure  []float32
}

func CreateApplication(name string) *Application {
	application := new(Application)
	if createdWindow, err := window.InitWindow(width, height, name); err != nil {
		panic(err)
	} else {
		application.window = createdWindow
	}

	return application
}

func (app *Application) Run() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	app.VAO = MakeVao(app.Figure)

	for !app.window.ShouldClose() {
		app.draw()
	}

}

func (app *Application) draw() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.UseProgram(app.Program)
	gl.BindVertexArray(app.VAO)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(app.Figure)/3))

	glfw.PollEvents()
	app.window.SwapBuffers()
}

func MakeVao(points []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	return vao
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
