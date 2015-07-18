package component

import (
    "fmt"
)

type Or2Recognizer struct {
    blueprint [][]string
}

func (o Or2Recognizer) Blueprint() [][]string {
    return o.blueprint
}

func (o Or2Recognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    ina, inb, out := 0, 0, 0
    connected := false
    width, height := len(o.blueprint) - 1, len(o.blueprint[0]) - 1
    or2com := Or2Component {id:id, X:x, Y:y, Width:width, Height:height, Connected:&connected, InA:&ina, InB:&inb, Out:&out}
    return &or2com
}

func NewOr2Recognizer() Or2Recognizer {
    blueprint := [][]string {{"-", "-", "-", ")", ""},
                             {"|", " ", ".", " ", ">"},
                             {"-", "-", "-", ")", ""}}
    or2rec := Or2Recognizer {blueprint:blueprint}
    return or2rec
}

type Or2Component struct {
    id string
    X, Y, Width, Height int
    Connected *bool
    InA *int
    InB *int
    Out *int
}

func (o Or2Component) Id() string {
    return o.id
}

func (o Or2Component) Space() (int, int, int, int) {
    return o.X, o.Y, o.Width, o.Height
}

func (o Or2Component) Update() {
    if *o.InA == 1 || *o.InB == 1 {
        *o.Out = 1
    } else {
        *o.Out = 0
    }
}

func (o Or2Component) Print() {
    fmt.Println("Or2 Gate ID:", o.id, "InA:", *o.InA, "InB:", *o.InB, "Out:", *o.Out)
}

func (o *Or2Component) Connect(c *int, t string) {
    if *o.Connected {
        o.InB = c
    } else {
        *o.Connected = true
        o.InA = c
    }
}

func (o Or2Component) Output(t string) *int {
    return o.Out
}

func (o Or2Component) InputStreams() []string {
    return []string {"I"}
}

func (o Or2Component) OutputStreams() []string {
    return []string {"O"}
}






