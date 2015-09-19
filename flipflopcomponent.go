package textelectronics

import "fmt"

type FlipFlopRecognizer struct {
    blueprint [][]string
}

func (f FlipFlopRecognizer) Blueprint() [][]string {
    return f.blueprint
}

func (f FlipFlopRecognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    data, last, clock, out, nout := 0, 0, 0, 0, 1
    width, height := len(f.blueprint) - 1, len(f.blueprint[0])
    flipflopcom := FlipFlopComponent {id:id, X:x, Y:y, Width:width, Height:height, last:&last, Data:&data, Clock:&clock, Out:&out, Nout:&nout}
    return &flipflopcom
}

func NewFlipFlopRecognizer() FlipFlopRecognizer {
    blueprint := [][]string {{"-", "-", "-", "-", "-"},
                             {"|", " ", ".", " ", "|"},
                             {"|", " ", " ", " ", "|"},
                             {"-", "-", "-", "-", "-"}}
    flipfloprec := FlipFlopRecognizer {blueprint:blueprint}
    return flipfloprec
}

type FlipFlopComponent struct {
    id string
    X, Y, Width, Height int
    last *int
    Data *int
    Clock *int
    Out *int
    Nout *int
}

func (f FlipFlopComponent) Id() string {
    return f.id
}

func (f FlipFlopComponent) Space() (int, int, int, int) {
    return f.X, f.Y, f.Width, f.Height
}

func (f FlipFlopComponent) Update() {
    if *f.Clock == 1 && *f.last == 0 {
        *f.Out = *f.Data
        if *f.Data == 0 {
            *f.Nout = 1
        } else {
            *f.Nout = 0
        }
    }
    *f.last = *f.Clock
}

func (f FlipFlopComponent) Print() {
    fmt.Println("Flip Flop ID:", f.id, "Data:", *f.Data, "Clock:", *f.Clock, "Out:", *f.Out, "NOut:", *f.Nout)
}

func (f *FlipFlopComponent) Connect(c *int, t string) {
    if t == "D" {
        f.Data = c
    }
    if t == "C" {
        f.Clock = c
    }
}

func (f FlipFlopComponent) Output(t string) *int {
    if t == "O" {
        return f.Out
    }
    if t == "Q" {
        return f.Nout
    }
    return nil
}

func (f FlipFlopComponent) InputStreams() []string {
    return []string {"D", "C"}
}

func (f FlipFlopComponent) OutputStreams() []string {
    return []string {"O", "Q"}
}

func (f FlipFlopComponent) Visual() map[Coordinate]*string {
    return nil
}

