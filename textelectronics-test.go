package main

import (

    "github.com/furryfaust/textelectronics/circuit"
    "github.com/furryfaust/textelectronics/components"
)

func main() {
    c := circuit.NewCircuit()

    c.AddRecognizer(components.NewAnd2Recognizer())
    c.AddRecognizer(components.NewInputRecognizer())
    c.AddRecognizer(components.NewProbeRecognizer())
    c.AddRecognizer(components.NewOr2Recognizer())
    c.AddRecognizer(components.NewNotRecognizer())
    c.AddRecognizer(components.NewXor2Recognizer())
    c.AddRecognizer(components.NewClockRecognizer())
    c.AddRecognizer(components.NewFlipFlopRecognizer())
    c.AddRecognizer(components.NewHexaconvRecognizer())

    c.Run(`example_circuits/circuit_four.txt`)
}
