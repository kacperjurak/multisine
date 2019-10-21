package main

import (
	"flag"
	"fmt"
	"strconv"
)
import "multisine"

type arrayFlags []float32

func (i *arrayFlags) String() string {
	return fmt.Sprintf("%g", *i)
}

func (i *arrayFlags) Set(value string) error {
	tmp, err := strconv.ParseFloat(value, 32)
	if err != nil {
		*i = append(*i, -1)
	} else {
		*i = append(*i, float32(tmp))
	}
	return nil
}

var (
	sf       uint
	duration float64
	freqs    arrayFlags
	amps     arrayFlags
	phases   arrayFlags
	sines    []multisine.Sine
	output   string
)

func main() {

	flag.UintVar(&sf, "sf", 1000, "Sampling frequency (Hz)")
	flag.Float64Var(&duration, "d", 1, "Signal duration (seconds)")
	flag.Var(&freqs, "f", "Sine frequency (Hz)")
	flag.Var(&amps, "a", "Sine amplitude")
	flag.Var(&phases, "p", "Sine phase (angle)")
	flag.Parse()
	if len(freqs) != len(amps) || len(freqs) != len(phases) || len(amps) != len(phases) {
		panic("Number of sines parameters are not equal")
	}
	sg := multisine.NewSignalGenerator(uint32(sf), float32(duration))

	for i, freq := range freqs {
		sines = append(sines, sg.GenerateSine(freq, amps[i], phases[i]))
	}

	ms := sg.GenerateMultisine(sines)

	for _, i := range ms {
		output += fmt.Sprintf("%g\r\n", i)
	}
	fmt.Print(output)
}
