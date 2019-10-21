package multisine

import (
	"math"
)

type Sine []float32
type Multisine []float32

type SignalGenerator struct {
	SamplingFrequency uint32
	Duration          float32
}

func NewSignalGenerator(f uint32, d float32) *SignalGenerator {
	return &SignalGenerator{
		f,
		d,
	}
}

// GenerateSine returns a slice with sine samples
func (sg *SignalGenerator) GenerateSine(frequency float32, amplitude float32, phase float32) Sine {
	var (
		samples          = sg.getNumberOfSamples()
		radiansPerSample = float64(frequency * 2 * math.Pi / float32(sg.SamplingFrequency))
		waveform         = make([]float32, samples)
		actPhase         = float64(phase * math.Pi / 180)
	)
	for i := uint32(0); i < samples; i++ {
		v := amplitude * float32(math.Sin(actPhase))
		waveform[i] = v
		actPhase += radiansPerSample
	}
	return waveform
}

// GenerateSine returns a slice with a summed sine samples (multisine)
func (sg *SignalGenerator) GenerateMultisine(sines []Sine) Multisine {
	sum := sg.createMultisine()
	for _, sine := range sines {
		for i, v := range sine {
			sum[i] = + v
		}
	}
	return sum
}

// getNumberOfSamples returns a number of samples in a signal
func (sg *SignalGenerator) getNumberOfSamples() uint32 {
	return uint32(float32(sg.SamplingFrequency) * sg.Duration)
}

// createMultisine creates an empty slice with length of a signal
func (sg *SignalGenerator) createMultisine() Multisine {
	return make([]float32, sg.getNumberOfSamples())
}
