package textelectronics

import (
    "fmt"
    "strconv"
)

type ClockRecognizer struct {
    blueprint [][]string
}

func (c ClockRecognizer) Blueprint() [][]string {
    return c.blueprint
}

func (c ClockRecognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    in, out := 1, 0
    tick := 0
    interval, err := strconv.Atoi(input[id])
    if err != nil {
        interval = 1
    }
    width, height := len(c.blueprint) - 1, len(c.blueprint[0]) - 1
    clockcom := ClockComponent {id:id, X:x, Y:y, Width:width, Height:height, Tick:tick, Interval:interval, In:&in, Out:&out}
    return &clockcom
}  

func NewClockRecognizer() ClockRecognizer {
    blueprint := [][]string {{" ", "*", " "},
                             {"*", ".", "*"},
                             {" ", "*", " "}}
    clockrec := ClockRecognizer {blueprint:blueprint}
    return clockrec
}

type ClockComponent struct {
    id string
    X, Y, Width, Height int
    Tick, Interval int
    In *int
    Out *int
}

func (c ClockComponent) Id() string {
    return c.id
}

func (c ClockComponent) Space() (int, int, int, int) {
    return c.X, c.Y, c.Width, c.Height
}

func (c *ClockComponent) Update() {
    if *c.In == 0 {
        c.Tick++
        if c.Tick == c.Interval {
            c.Tick = 0
            if *c.Out == 0 {
                *c.Out = 1
            } else {
                *c.Out = 0
            }
        }
    }
}

func (c ClockComponent) Print() {
    fmt.Println("Clock ID:", c.id, "Out:", *c.Out)
}

func (cl *ClockComponent) Connect(c *int, t string) {
    cl.In = c
}

func (c ClockComponent) Output(t string) *int {
    return c.Out
}

func (c ClockComponent) InputStreams() []string {
    return []string {"I"}
}

func (c ClockComponent) OutputStreams() []string {
    return []string {"O"}
}

func (c ClockComponent) Visual() map[Coordinate]*string {
    return nil
}
