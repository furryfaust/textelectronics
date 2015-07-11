package component

import (
    "fmt"
)

type And2Recognizer struct {
    blueprint [][]string
}

func (a And2Recognizer) Blueprint() [][]string {
    return a.blueprint
}

func (a And2Recognizer) NewComponent(id string, x int, y int, width int, height int) Component{
    ina, inb, out := 0, 0, 0
    and2com := And2Component {id:id, X:x, Y:y, Width:width, Height:height, InA:&ina, InB:&inb, Out:&out}
    return and2com
}

func NewAnd2Recognizer() And2Recognizer {
    blueprint := [][]string {{"-", "-", "-", "\\", ""},
                             {"|", " ", ".", " ", ">"},
                             {"-", "-", "-", "/", ""}}
    and2rec := And2Recognizer {blueprint:blueprint}
    return and2rec
}

type And2Component struct {
    id string
    X, Y, Width, Height int
    Connected *bool
    InA *int
    InB *int
    Out *int
}

func (a And2Component) Id() string {
    return a.id
}

func (a And2Component) Space() (int, int, int, int) {
    return a.X, a.Y, a.Width, a.Height
}

func (a And2Component) Update() {
    if *a.InA == 0 || *a.InB == 0 {
        *a.Out = 0
    } else {
        *a.Out = 1
    }
}

func (a And2Component) Print() {
    fmt.Println("InA:", a.InA, "InB:", a.InB, "Out", a.Out)
}

func (a And2Component) Input(t string) *int {
    if t == "I" {
        if *a.Connected {
            return a.InB
        } else {
            *a.Connected = true
            return a.InA
        }
    }
    return nil
}

func (a And2Component) Output(t string) *int {
    return a.Out
}


