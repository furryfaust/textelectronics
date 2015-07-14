package circuitutils

import (
    "os"
    "fmt"
    "bufio"
    "regexp"
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

    
    getComponentById := func(id string) component.Component {
        for i := range components {
            if components[i].Id() == id {
                return components[i]
            }
        }
        return nil
    }

    getComponentByLocation := func(x int, y int) component.Component {
        for i := range components {
            cX, cY, cWidth, cHeight := components[i].Space()
            if x >= cX && x <= cX + cWidth && y >= cY && y <= cY + cHeight {
                return components[i]
            }
        }
        return nil
    }

    contains := func(query string, slice []string) bool {
        for index := range slice {
            if slice[index] == query {
                return true
            }
        }
        return false
    }

    recognizeIOType := func(x int, y int) (string, string, int, int, int) {
        if x > 0 && len(rawc) - 1 > x {
            com := getComponentByLocation(x - 1, y)
            if com != nil {
                if contains(rawc[x + 1][y], com.OutputStreams()) || contains(rawc[x + 1][y], com.InputStreams()) {
                    return rawc[x + 1][y], com.Id(), x + 2, y, 0
                }
            }
        }
        if x < len(rawc) && x > 0 {
            com := getComponentByLocation(x + 1, y)
            if com != nil {
                if contains(rawc[x - 1][y], com.OutputStreams()) || contains(rawc[x - 1][y], com.InputStreams()) {
                    return rawc[x - 1][y], com.Id(), x - 2, y, 1
                }
            }
        }
        if y < len(rawc[0]) {
            com := getComponentByLocation(x, y - 1)
            if com != nil {
                if contains(rawc[x][y + 1], com.OutputStreams()) || contains(rawc[x][y + 1], com.InputStreams()) {
                    return rawc[x][y + 1], com.Id(), x, y + 2, 2
                }
            }
        }
        if y > 0 {
            com := getComponentByLocation(x, y + 1)
            if com != nil {
                if contains(rawc[x][y - 1], com.OutputStreams()) || contains(rawc[x][y - 1], com.InputStreams()) {
                    return rawc[x][y - 1], com.Id(), x, y - 2, 3
                }
            }
        }
        return "", "", 0, 0, 0
    }

    for y := 0; y != len(rawc[0]); y++ {
        for x := 0; x != len(rawc); x++ {
            if rawc[x][y] == "%" {
                mio, mid, cX, cY, direction := recognizeIOType(x, y)
                    mcom := getComponentById(mid)

                    if contains(mio, mcom.OutputStreams()) {
                        var seekInput func(x int, y int, direction int)
                        seekInput = func(x int, y int, direction int) {
                            switch direction {
                                case 0:
                                    x++
                                    if match, _ := regexp.MatchString("[|a-zA-Z]", rawc[x][y]); match {
                                        seekInput(x, y, direction)
                                    }
                                    if rawc[x][y] == "%" {
                                        fio, fid, _, _, _ := recognizeIOType(x, y)

                                        fcom := getComponentById(fid)
                                        fmt.Println(mid, "connected to", fid)
                                        *fcom.Input(fio) = *mcom.Output(mio)
                                        fcom.Update()
                                    }
                                    if rawc[x][y] == "+" {
                                        seekInput(x, y, 2)
                                        seekInput(x, y, 3)
                                    }
                                case 1:
                                    x--
                                    if match, _ := regexp.MatchString("[|a-zA-Z]", rawc[x][y]); match {
                                        seekInput(x, y, direction)
                                    }
                                    if rawc[x][y] == "%" {
                                        fio, fid, _, _, _ := recognizeIOType(x, y)

                                        fcom := getComponentById(fid)
                                        fmt.Println(mid, "connected to", fid)
                                        *fcom.Input(fio) = *mcom.Output(mio)
                                        fcom.Update()
                                    }
                                    if rawc[x][y] == "+" {
                                        seekInput(x, y, 2)
                                        seekInput(x, y, 3)
                                    }
                                case 2:
                                    y++
                                    if match, _ := regexp.MatchString("[-a-zA-Z]", rawc[x][y]); match {
                                        seekInput(x, y, direction)
                                    }
                                    if rawc[x][y] == "%" {
                                        fio, fid, _, _, _ := recognizeIOType(x, y)

                                        fcom := getComponentById(fid)
                                        fmt.Println(mid, "connected to", fid)
                                        *fcom.Input(fio) = *mcom.Output(mio)
                                        fcom.Update()
                                    }
                                    if rawc[x][y] == "+" {
                                        seekInput(x, y, 0)
                                        seekInput(x, y, 1)
                                    }
                                case 3:
                                    y--
                                    if match, _ := regexp.MatchString("[-a-zA-Z]", rawc[x][y]); match {
                                        seekInput(x, y, direction)
                                    }
                                    if rawc[x][y] == "%" {
                                        fio, fid, _, _, _ := recognizeIOType(x, y)

                                        fcom := getComponentById(fid)
                                        fmt.Println(mid, "connected to", fid)
                                        *fcom.Input(fio) = *mcom.Output(mio)
                                        fcom.Update()
                                    }
                                    if rawc[x][y] == "+" {
                                        seekInput(x, y, 0)
                                        seekInput(x, y, 1)
                                    }
                            }}
                        seekInput(cX, cY, direction)
                    }
            }
        }
    }

    for index := range *c.Components {
        (*c.Components)[index].Print()
    }
}
















