This is a package for generating sine signals and sum sines into one multisine signal.

#### Generate a single sine

Create a new SignalGenerator.

sampling frequency: 44.1 kHz, signal duration: 2 seconds

```go
sg := multisine.NewSignalGenerator(44100, 2)
```

Generate a sine.

sine frequency: 1 kHz, amplitude: 2 (no unit), phase: 180 deg
```go
sg.GenerateSine(1000, 2, 180)
```

#### Generate a multisine

Create a SignalGenerator and array of sines then generate a one multisine.
```go
var sines []multisine.Sine = make([]multisine.Sine, 3)

sg := multisine.NewSignalGenerator(65536, 1)

sines[0] = sg.GenerateSine(1000, 1, 0)
sines[1] = sg.GenerateSine(100, 0.5, 90)
sines[2] = sg.GenerateSine(10, 0.7, 270)

sg.GenerateMultisine(sines)
```

#### CLI

Use a ```cmd/multisine``` command to generate single sine or multisine.

##### Single sine
```
./multisine -sf 44100 -d 2 -f 1000 -a 2 -p 180 > sine.txt
```
##### Multisine
```
./multisine -sf 65536 -d 1 -f 1000 -a 1 -p 0 -f 100 -a 0.5 -p 90 -f 10 -a 0.7 -p 270 > multisine.txt
```