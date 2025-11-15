package visual

import (
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
)

type Wave struct {
	Vertices []float32
	Samples  int32

	phase float32
	vao   uint32

	posUniform, hzUniform int32
}

func NewWave(samples int32, vao uint32, posUniform, hzUniform int32) *Wave {
	v := make([]float32, samples)
	for i := range samples {
		v[i] = (float32(float64(i)/float64((samples-1)))*2.0 - 1.0)
	}
	return &Wave{
		Vertices:   v,
		Samples:    samples,
		vao:        vao,
		posUniform: posUniform,
		hzUniform:  hzUniform,
	}
}

func (w *Wave) Draw(program uint32) {
	w.phase += float32(float64(time.Second) / 1e9)

	gl.UseProgram(program)
	// When we have data, we can just pass it through here.
	gl.Uniform1f(w.posUniform, w.phase)
	gl.Uniform1f(w.hzUniform, 60)

	gl.BindVertexArray(w.vao)

	gl.ActiveTexture(gl.TEXTURE0)

	gl.DrawArrays(gl.LINE_STRIP, 0, w.Samples)
}
