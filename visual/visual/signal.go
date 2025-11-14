package visual

import "math"

type Wave struct {
	Vertices []float32
	Rate     int64
}

func NewWave(samples int64) *Wave {
	v := make([]float32, samples*2)
	for i := range samples {
		v[i*2] = (float32(float64(i)/float64((samples-1)))*2.0 - 1.0)
		v[i*2+1] = float32(math.Sin(float64(i) / float64(samples-1) * math.Pi * 2))
	}
	return &Wave{
		Vertices: v,
		Rate:     samples,
	}
}
