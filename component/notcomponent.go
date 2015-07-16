package component

import (
    "fmt"
)

type NotRecognizer struct {
    blueprint [][]string
}

func (n NotRecognizer) Blueprint() [][]string {
    return n.blueprint
}

func (n NotRecognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    in, out := 0, 1
    width, height := len(n.blueprint) - 1, len(n.blueprint[0]) - 1
    notcom := NotComponent {id:id, X:x, Y:y, Width:width, Height:height, In:&in, Out:&out}
    return notcom
}

func NewNotRecognizer() NotRecognizer {
    blueprint := [][]string {{"!", ".", "!", ""}}
    notrec := NotRecognizer {blueprint:blueprint}
    return notrec
}

type NotComponent struct {
    id string
    X, Y, Width, Height int
    In *int
    Out *int
}

func (n NotComponent) Id() string {
    return n.id
}

func (n NotComponent) Space() (int, int, int, int) {
    return n.X, n.Y, n.Width, n.Height
}

func (n NotComponent) Update() {
    if *n.In == 0 {
        *n.Out = 1
    } else {
        *n.Out = 0
    }
}

func (n NotComponent) Print() {
    fmt.Println("Not Gate ID:", n.id, "In:", *n.In, "Out:", *n.Out)
}

func (n NotComponent) Input(t string) *int {
    return n.In
}

func (n NotComponent) Output(t string) *int {
    return n.Out
}

func (n NotComponent) InputStreams() []string {
    return []string {"I"}
}

func (n NotComponent) OutputStreams() []string {
    return []string {"O"}
}





