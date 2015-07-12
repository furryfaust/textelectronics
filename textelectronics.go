package main

import (    
    "os"
    "github.com/furryfaust/textelectronics/component"
    "github.com/furryfaust/textelectronics/circuitutils"
)

func main() {
    circuit := circuitutils.NewCircuit()
    circuit.AddRecognizer(component.NewAnd2Recognizer())

    circuit.Parse(os.Args[1])
}