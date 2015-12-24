package components

import (
    "fmt"
    "strconv"
)

type HexaconvRecognizer struct {
    blueprint [][]string
}

func (h HexaconvRecognizer) Blueprint() [][]string {
    return h.blueprint
}

func (h HexaconvRecognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
    ina, inb, inc, ind, out := 0, 0, 0, 0, "0"
    visual := map[Coordinate]*string { Coordinate { X:2, Y:2 }:&out }
    width, height := len(h.blueprint) - 1, len(h.blueprint[0]) - 1
    hexaconvcom := HexaconvComponent {id:id, X:x, Y:y, Width:width, Height:height, InA:&ina, InB:&inb, InC:&inc, InD:&ind, visual:visual}
    return &hexaconvcom
}

func NewHexaconvRecognizer() HexaconvRecognizer {
    blueprint := [][]string{{"-", "-", "-", "-", "-", "-"},
                            {"|", " ", " ", " ", " ", "|"},
                            {"|", " ", ".", " ", " ", "|"},
                            {"|", " ", " ", " ", " ", "|"},
                            {"-", "-", "-", "-", "-", "-"}}
    hexaconvrec := HexaconvRecognizer{blueprint: blueprint}
    return hexaconvrec
}

type HexaconvComponent struct {
    id string
    X, Y, Width, Height int
    InA, InB, InC, InD *int
    Out *string
    visual map[Coordinate]*string
}

func (h HexaconvComponent) Id() string {
    return h.id
}

func (h HexaconvComponent) Space() (int, int, int, int) {
    return h.X, h.Y, h.Width, h.Height
}

func (h HexaconvComponent) Update() {
    sum := 0
    out := h.visual[Coordinate {X: 2, Y:2}]
    if *h.InA == 1 {
        sum += 1
    }
    if *h.InB == 1 {
        sum += 2
    }
    if *h.InC == 1 {
        sum += 4
    }
    if *h.InD == 1 {
        sum += 8
    }
    if sum < 10 {
       *out = strconv.Itoa(sum)
    } else {
        if sum == 10 {
            *out = "A"
        } else if sum == 11 {
            *out = "B"
        } else if sum == 12 {
            *out = "C"
        } else if sum == 13 {
            *out = "D"
        } else if sum == 14 {
            *out = "E"
        } else if sum == 15 {
            *out = "F"
        }
    }
}

func (h HexaconvComponent) Print() {
    fmt.Println("Hexaconv ID:", h.id, "InA:", *h.InA, "InB:", *h.InB, "InC:", *h.InC, "InD:", *h.InD)
}

func (h *HexaconvComponent) Connect(c *int, t string) {
    if t == "A" {
       h.InA = c
    } else if t == "B" {
        h.InB = c
    } else if t == "C" {
        h.InC = c
    } else if t == "D" {
        h.InD = c
    }
}

func (h HexaconvComponent) Output(t string) *int {
    return nil
}

func (h HexaconvComponent) InputStreams() []string {
    return []string {"A", "B", "C", "D"}
}

func (h HexaconvComponent) OutputStreams() []string {
    return []string {}
}

func (h HexaconvComponent) Visual() map[Coordinate]*string {
    return h.visual
}
