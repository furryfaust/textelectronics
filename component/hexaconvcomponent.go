package component

import "fmt"

type HexaconvRecognizer struct {
    blueprint [][]string
}

func (h HexaconvRecognizer) Blueprint() [][]string {
    return h.blueprint
}

func (h HexaconvRecognizer) NewComponent(id string, x int, y int, input map[string]string) Component {
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
}

func (h HexaconvComponent) Id() string {
    return h.id
}

func (h HexaconvComponent) Space() (int, int, int, int) {
    return h.X, h.Y, h.Width, h.Height
}

func (h HexaconvComponent) Update() {
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
 
