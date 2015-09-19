package textelectronics

import "fmt"

type Xor2Recognizer struct {
    blueprint [][]string
}

func (x Xor2Recognizer) Blueprint() [][]string {
    return x.blueprint
}

func (xo Xor2Recognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    ina, inb, out := 0, 0, 0
    connected := false
    width, height := len(xo.blueprint) - 1, len(xo.blueprint[0]) - 1
    xor2com := Xor2Component {id:id, X:x, Y:y, Width:width, Height:height, Connected:&connected, InA:&ina, InB:&inb, Out:&out}
    return &xor2com
}

func NewXor2Recognizer() Xor2Recognizer {
    blueprint := [][]string {{"~", "~", "~", ")", ""},
                             {"|", " ", ".", " ", ">"},
                             {"~", "~", "~", ")", ""}}
    xor2rec := Xor2Recognizer {blueprint:blueprint}
    return xor2rec
}

type Xor2Component struct {
    id string
    X, Y, Width, Height int
    Connected *bool
    InA *int
    InB *int
    Out *int
}

func (x Xor2Component) Id() string {
    return x.id
}

func (x Xor2Component) Space() (int, int, int, int) {
    return x.X, x.Y, x.Width, x.Height
}

func (x Xor2Component) Update() {
    if *x.InA != *x.InB {
        *x.Out = 1
    } else {
        *x.Out = 0
    }
}

func (x Xor2Component) Print() {
    fmt.Println("Xor2 Gate ID:", x.id, "InA:", *x.InA, "InB:", *x.InB, "Out:", *x.Out)
}

func (x *Xor2Component) Connect(c *int, t string) {
    if *x.Connected {
        x.InB = c
    } else {
        *x.Connected = true
        x.InA = c
    }
}

func (x Xor2Component) Output(t string) *int {
    return x.Out
}

func (x Xor2Component) InputStreams() []string {
    return []string {"I"}
}

func (x Xor2Component) OutputStreams() []string {
    return []string {"O"}
}

func (x Xor2Component) Visual() map[Coordinate]*string {
    return nil
}
