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


