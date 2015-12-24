package components

import (
    "fmt"
    "strconv"
)

type InputRecognizer struct {
    blueprint [][]string
}

func (i InputRecognizer) Blueprint() [][]string {
    return i.blueprint
}

func (i InputRecognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    rawo := input[id]
    out, err := strconv.Atoi(rawo)
    if err != nil {
        out = 0
    }
    width, height := len(i.blueprint) - 1, len(i.blueprint[0]) - 1
    inputcom := InputComponent {id:id, X:x, Y:y, Width:width, Height:height, Out:&out}
    return &inputcom
}

func NewInputRecognizer() InputRecognizer {
    blueprint := [][]string {{"{", ".", "}"}}
    inputrec := InputRecognizer {blueprint:blueprint}
    return inputrec
}

type InputComponent struct {
    id string
    X, Y, Width, Height int
    Out *int
}

func (i InputComponent) Id() string {
    return i.id
}

func (i InputComponent) Space() (int, int, int, int) {
    return i.X, i.Y, i.Width, i.Height
}

func (i InputComponent) Update() {}

func (i InputComponent) Print() {
    fmt.Println("Input ID:", i.id, "Out:", *i.Out)
}

func (i *InputComponent) Connect(c *int, t string) {}

func (i InputComponent) Output(t string) *int {
    return i.Out
}

func (i InputComponent) InputStreams() []string {
    return []string {}
}

func (i InputComponent) OutputStreams() []string {
    return []string {"O"}
}

func (i InputComponent) Visual() map[Coordinate]*string {
    return nil
}
