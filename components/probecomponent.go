package components

import (
    "fmt"
    "strconv"
)

type ProbeRecognizer struct {
    blueprint [][]string
}

func (p ProbeRecognizer) Blueprint() [][]string {
    return p.blueprint
}

func (p ProbeRecognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    in, out, vout := 0, 0, "0"
    width, height := len(p.blueprint) - 1, len(p.blueprint[0]) - 1
    visual := map[Coordinate]*string { Coordinate {X:0, Y:1}:&vout } 
    probecom := ProbeComponent {id:id, X:x, Y:y, Width:width, Height:height, In:&in, Out:&out, visual:visual}
    return &probecom
}

func NewProbeRecognizer() ProbeRecognizer {
    blueprint := [][]string {{"(", ".", ")"}}
    proberec := ProbeRecognizer {blueprint:blueprint}
    return proberec
}

type ProbeComponent struct {
    id string
    X, Y, Width, Height int
    In *int
    Out *int
    visual map[Coordinate]*string
}

func (p ProbeComponent) Id() string {
    return p.id
}

func (p ProbeComponent) Space() (int, int, int, int) {
    return p.X, p.Y, p.Width, p.Height
}

func (p ProbeComponent) Update() {
    *p.Out = *p.In
    *p.visual[Coordinate {X:0, Y:1}] = strconv.Itoa(*p.Out) 
}

func (p ProbeComponent) Print() {
    fmt.Println("Probe ID:", p.id, "In:", *p.In, "Out:", *p.Out)
}

func (p *ProbeComponent) Connect(c *int, t string) {
    p.In = c
}

func (p ProbeComponent) Output(t string) *int {
    return p.Out
}

func (p ProbeComponent) InputStreams() []string {
    return []string {"I"}
}

func (p ProbeComponent) OutputStreams() []string {
    return []string {"O"}
}

func (p ProbeComponent) Visual() map[Coordinate]*string {
    return p.visual
}
