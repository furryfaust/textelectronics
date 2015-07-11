package component

type And2Recognizer struct {
    Blueprint [][]string
}

func (a And2Recognizer) Blueprint() [][]string {
    return a.Blueprint
}

func NewAnd2Recognizer() And2Recognizer {
    blueprint := make([][]string) { {"-", "-", "-", "\\", ""}
                                    {"|", " ", ".", " ", ">"}
                                    {"-", "-", "-", "/", ""}}
    and2rec := And2Recognizer {Blueprint:blueprint}
    return and2rec
}

