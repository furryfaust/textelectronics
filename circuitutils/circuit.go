package circuitutils

import (
    "os"
    "bufio"
    "strings"
    "io/ioutil"
    "github.com/furryfaust/textelectronics/component"
)

type Circuit struct {
    Recognizers *[]component.Recognizer
    Components *[]component.Component
}

func NewCircuit() Circuit {
    circuit := Circuit {}
    recognizers := make([]component.Recognizer, 0)
    circuit.Recognizers = &recognizers
    components := make([]component.Component, 0)
    circuit.Components = &components
    return circuit
}

func (c Circuit) AddRecognizer(recognizer component.Recognizer) {
    recognizers := *c.Recognizers
    *c.Recognizers = append(recognizers, recognizer)
}

func (c Circuit) Parse(path string) {
    dat, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }
    lines := strings.Count(string(dat), "\n")

    file, err := os.Open(path)
    if err != nil {
        panic(err)
    }

    longest := 0
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        if len(scanner.Text()) > longest {
            longest = len(scanner.Text())
        }
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    rawc := make([][]string, lines + 1)
    for key := range rawc {
        rawc[key] = make([]string, longest + 1)
    }

    file.Seek(0, 0)
    scanner = bufio.NewScanner(file)
    i := 0
    for scanner.Scan() {
        for j, k := range []byte(scanner.Text()) {
            rawc[i][j] = string(k)
        }
        i++
    }

    if err := scanner.Err(); err != nil {
        panic(err)
    }

    components := *c.Components

    recognizeComponent := func(x int, y int, blueprint [][]string) (bool, string) {
        var id string
        if (rawc[x][y] == blueprint[0][0]) {
            for i := 0; i != len(blueprint); i++ {
                for j := 0; j != len(blueprint[0]); j++ {
                    if rawc[x + i][y + j] != blueprint[i][j] && blueprint[i][j] != "." && blueprint[i][j] != "" {
                        return false, id
                    } else if blueprint[i][j] == "." {
                        id = rawc[x + i][y + j]
                    }
                }
            }
            return true, id
        }
        return false, id
    }

    values := make(map[string]string)
    if len(os.Args) > 2 {
        for i := 2; i != len(os.Args); i++ {
            splt := strings.Split(os.Args[i], ":")
            values[splt[0]] = splt[1]
        }
    }

    for y := 0; y != len(rawc[0]); y++ {
        for x := 0; x != len(rawc); x++ {
            for index := range *c.Recognizers {
                rec := (*c.Recognizers)[index]
                if found, id := recognizeComponent(x, y, rec.Blueprint()); found {
                    com := rec.NewComponent(id, x, y, values)
                    components = append(components, com)
                    *c.Components = components
                }
            }
        }
    }

    getComponentById := func(id string) *Component {
        
    }

    for index := range *c.Components {
        (*c.Components)[index].Print()
    }
}
















