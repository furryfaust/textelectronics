package circuitutils

import (
    "os"
    "fmt"
    "github.com/furryfaust/textelectronics/component"
)

type Circuit struct {
    Recognizers *[]component.Recognizer
    Components *[]component.Component
}

func NewCircuit() Circuit {
    circuit := Circuit {}
    recognizers := make([]component.Recognizer, 0)
    circuit.Recognizers = &recognizers

    fmt.Println(*circuit.Recognizers)
    return circuit
}

func (c Circuit) AddRecognizer(recognizer Recognizer) {
    recognizers := *c.Recognizers
    recognizers = append(recognizers, recognizer)
}

func (c Circuit) Parse(path string) {

}

