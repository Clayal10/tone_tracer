package visual

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Wave struct {
	Vertices []float32
	Samples  int32

	vao uint32

	index int

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

var Values = []float32{
	0.01, 0.08, 0.13, 0.20, 0.25, 0.21, 0.14, 0.08, 0.01, -0.04, -0.1, -0.18, -0.2, -0.16, -0.11, -0.07, -0.01,
}

// NOTES:
//   - ideally, only the right most value will be updated with the new value and all others will just
//     'shift' to the left
//   - Would essentially need to know the next (to the right) value.
func (w *Wave) Draw(program uint32) {
	length := len(Values)
	y := Values[w.index%length]
	w.index++

	gl.UseProgram(program)
	// When we have data, we can just pass it through here.
	gl.Uniform1f(w.posUniform, y)
	gl.Uniform1f(w.hzUniform, 60)

	gl.BindVertexArray(w.vao)

	gl.ActiveTexture(gl.TEXTURE0)

	gl.DrawArrays(gl.LINE_STRIP, 0, w.Samples)
}
