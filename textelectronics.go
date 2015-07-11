package main

import (    
    "fmt"
    "github.com/furryfaust/textelectronics/component"
    "github.com/furryfaust/textelectronics/circuitutils"
)

func main() {
    circuit := circuitutils.NewCircuit()
    circuit.AddRecognizer(component.NewAnd2Recognizer())
    fmt.Println(circuit)
}