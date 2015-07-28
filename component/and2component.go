package component

import "fmt"

type And2Recognizer struct {
    blueprint [][]string
}

func (a And2Recognizer) Blueprint() [][]string {
    return a.blueprint
}

func (a And2Recognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    ina, inb, out := 0, 0, 0
    connected := false
    width, height := len(a.blueprint) - 1, len(a.blueprint[0]) - 1
    and2com := And2Component {id:id, X:x, Y:y, Width:width, Height:height, Connected:&connected, InA:&ina, InB:&inb, Out:&out}
    return &and2com
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
    fmt.Println("And2 Gate ID:", a.id, "InA:", *a.InA, "InB:", *a.InB, "Out:", *a.Out)
}

func (a *And2Component) Connect(c *int, t string) {
    if *a.Connected {
        a.InB = c
    } else {
        *a.Connected = true
        a.InA = c
    }
}

func (a And2Component) Output(t string) *int {
    return a.Out
}

func (a And2Component) InputStreams() []string {
    return []string {"I"}
}

func (a And2Component) OutputStreams() []string {
    return []string {"O"}
}

func (a And2Component) Visual() map[Coordinate]*string {
    return nil
}
