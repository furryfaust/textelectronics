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
                            {"|", " ", " ", " ", " ", "|"},
                            {"|", " ", " ", " ", " ", "|"},
                            {"-", "-", "-", "-", "-", "-"}}
    hexaconvrec := HexaconvRecognizer{blueprint: blueprint}
    return hexaconvrec
}
