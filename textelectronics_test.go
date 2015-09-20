package textelectronics

import (
    "testing"
)

func TestCircuit(t *testing.T) {
    circuit := NewCircuit()

    circuit.AddRecognizer(NewAnd2Recognizer())
    circuit.AddRecognizer(NewInputRecognizer())
    circuit.AddRecognizer(NewProbeRecognizer())
    circuit.AddRecognizer(NewOr2Recognizer())
    circuit.AddRecognizer(NewNotRecognizer())
    circuit.AddRecognizer(NewXor2Recognizer())
    circuit.AddRecognizer(NewClockRecognizer())
    circuit.AddRecognizer(NewFlipFlopRecognizer())
    circuit.AddRecognizer(NewHexaconvRecognizer())

    circuit.Simulate(`example_circuits/circuit_two.txt`)
}
