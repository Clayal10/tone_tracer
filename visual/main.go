package main

import (
	_ "image/png"
	"log"
	"runtime"

	"github.com/Clayal10/tone_tracer/visual/visual"
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/go-gl/mathgl/mgl32"
)

const windowWidth = 800
const windowHeight = 600
const vertexShaderPath = "shaders/vertex_shader.glsl"
const fragmentShaderPath = "shaders/fragment_shader.glsl"
const numberOfPoints = 3000

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Tone Tracer", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()

	// Initialize Glow
	if err := gl.Init(); err != nil {
		panic(err)
	}

	program, err := setupProgram()
	if err != nil {
		panic(err)
	}

	gl.UseProgram(program)

	projection := mgl32.Ortho(-1, 1, -1, 1, -1, 1)
	projectionUniform := gl.GetUniformLocation(program, gl.Str("projection\x00"))
	gl.UniformMatrix4fv(projectionUniform, 1, false, &projection[0])

	textureUniform := gl.GetUniformLocation(program, gl.Str("tex\x00"))
	gl.Uniform1i(textureUniform, 0)

	gl.BindFragDataLocation(program, 0, gl.Str("outputColor\x00"))

	// Configure the vertex data
	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	posUniform := gl.GetUniformLocation(program, gl.Str("pos\x00"))
	gl.Uniform1f(posUniform, 0)
	hzUniform := gl.GetUniformLocation(program, gl.Str("hz\x00"))
	gl.Uniform1f(hzUniform, 0)

	wave := visual.NewWave(numberOfPoints, vao, posUniform, hzUniform)

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(wave.Vertices), gl.Ptr(wave.Vertices), gl.STATIC_DRAW)

	xAttrib := uint32(gl.GetAttribLocation(program, gl.Str("x\x00")))
	gl.EnableVertexAttribArray(xAttrib)
	gl.VertexAttribPointerWithOffset(xAttrib, 1, gl.FLOAT, false, 1*4, 0)

	// Configure global settings
	gl.Disable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(1.0, 1.0, 1.0, 1.0)

	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		wave.Draw(program)

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
