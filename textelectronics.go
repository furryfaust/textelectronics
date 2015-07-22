package main

import (    
    "os"
    "github.com/furryfaust/textelectronics/component"
    "github.com/furryfaust/textelectronics/circuitutils"
)

func main() {
    circuit := circuitutils.NewCircuit()
    circuit.AddRecognizer(component.NewAnd2Recognizer())
    circuit.AddRecognizer(component.NewInputRecognizer())
    circuit.AddRecognizer(component.NewProbeRecognizer())
    circuit.AddRecognizer(component.NewOr2Recognizer())
    circuit.AddRecognizer(component.NewNotRecognizer())
    circuit.AddRecognizer(component.NewXor2Recognizer())
    circuit.AddRecognizer(component.NewClockRecognizer())
    circuit.AddRecognizer(component.NewFlipFlopRecognizer())
    circuit.Simulate(os.Args[1])
}