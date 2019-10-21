package multisine_test

import (
	"multisine"
	"testing"
)

func TestNew(t *testing.T) {
	var (
		samplingFrequency uint32  = 44100
		duration          float32 = 1
	)
	sg := multisine.NewSignalGenerator(samplingFrequency, duration)
	if sg.SamplingFrequency != samplingFrequency {
		t.Errorf("Sampling Frequency doesn't match")
	}
	if sg.Duration != duration {
		t.Errorf("Duration doesn't match")
	}
}

func TestSineZeroPhase(t *testing.T) {
	var (
		samplingFrequency uint32  = 1000
		duration          float32 = 1
	)
	sg := multisine.NewSignalGenerator(samplingFrequency, duration)
	sine := sg.GenerateSine(1, 1, 0)
	if sine[0] != 0 || sine[250] != 1 || sine[500] != -6.8796325e-08 || sine[750] != -1 {
		t.Errorf("Wrong sine generated")
	}
}

func TestSineHalfPhase(t *testing.T) {
	var (
		samplingFrequency uint32  = 1000
		duration          float32 = 1
	)
	sg := multisine.NewSignalGenerator(samplingFrequency, duration)
	sine := sg.GenerateSine(1, 1, 180)
	if sine[0] != -8.742278e-08 || sine[250] != -1 || sine[500] != 1.562191e-07 || sine[750] != 1 {
		t.Errorf("Wrong sine generated")
	}
}

func TestMultisine(t *testing.T) {
	var (
		samplingFrequency uint32  = 10000
		duration          float32 = 1
		sines             []multisine.Sine
		freqs             = []float32{1000, 100, 10}
	)
	sg := multisine.NewSignalGenerator(samplingFrequency, duration)
	for _, freq := range freqs {
		sines = append(sines, sg.GenerateSine(freq, 1, 0))
	}
	ms := sg.GenerateMultisine(sines)
	if ms[0] != 0 || ms[250] != 1 || ms[500] != -6.8796325e-08 || ms[750] != -1 {
		t.Errorf("Wrong multisine generated")
	}
}
func TestMultisine1(t *testing.T) {
	var sines []multisine.Sine = make([]multisine.Sine, 3)
	sg := multisine.NewSignalGenerator(65536, 1)
	sines[0] = sg.GenerateSine(1000, 1, 0)
	sines[1] = sg.GenerateSine(100, 0.5, 90)
	sines[2] = sg.GenerateSine(10, 0.7, 270)
	sg.GenerateMultisine(sines)
}