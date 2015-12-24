package circuit

import (
//    "os"
    "log"
    "regexp"
    "strings"
    "io/ioutil"

    gc "github.com/rthornton128/goncurses"
    "github.com/furryfaust/textelectronics/components"
)

var (
    IO_CHECK = map[int]map[string]int {
        0: map[string]int {
            "com_x": -1, "com_y": 0, "type_x": 1, "type_y": 0, "x": 2, "y": 0,
        },
        1: map[string]int {
            "com_x": 1, "com_y": 0, "type_x": -1, "type_y": 0, "x": -2, "y": 0,
        },
        2: map[string]int {
            "com_x": 0, "com_y": -1, "type_x": 0, "type_y": 1, "x": 0, "y": 2,
        },
        3: map[string]int {
            "com_x": 0, "com_y": 1, "type_x": 0, "type_y": -1, "x": 0, "y": -2,
        },
    }
)

type Circuit struct {
    Recognizers []components.Recognizer
    Components []components.Component

    circuit [][]string
}

func NewCircuit() *Circuit {
    circuit := &Circuit {}
    recognizers := make([]components.Recognizer, 0)
    circuit.Recognizers = recognizers
    components := make([]components.Component, 0)
    circuit.Components = components
    return circuit
}

func (c *Circuit) AddRecognizer(recognizer components.Recognizer) {
    c.Recognizers = append(c.Recognizers, recognizer)
}

func (c *Circuit) Run(path string) {
    c.prepare(path)
    c.parse()
    c.assemble()

    c.display()
}

func (c *Circuit) prepare(path string) {
    dat, _ := ioutil.ReadFile(path)
    lines := strings.Split(string(dat), "\n")

    width := getLongestInSlice(lines)

    c.circuit = [][]string {}
    c.circuit = append(c.circuit, make([]string, width))
    for  _, line := range lines {
        chars := append([]string { " " }, strings.Split(line, "")...)
        chars = append(chars, make([]string, width - len(line) + 1)...)
        chars = append(chars, " ")
        c.circuit = append(c.circuit, chars)
    }
    c.circuit = append(c.circuit, make([]string, width))
}

func getLongestInSlice(slice []string) (longest int) {
    for _, s := range slice {
        if len(s) > longest {
            longest = len(s)
        }
    }
    return
}

func (c *Circuit) parse() {
    circuit := c.circuit
    values := map[string]string {}
    /*
    for i := 2; i != len(os.Args) ; i++ {
        pair := strings.Split(os.Args[i], ":")
        values[pair[0]] = pair[1]
    }*/

    for y := 0; y != len(circuit[0]); y++ {
        for x := 0; x != len(circuit); x++ {
            for _, rec := range c.Recognizers {
                if found, id := c.recognizeComponent(x, y, rec.Blueprint()); found && id != "" {
                    com := rec.NewComponent(id, x, y, values)
                    c.Components = append(c.Components, com)
                }
            }
        }
    }
}

func (c *Circuit) recognizeComponent(x int, y int, blueprint [][]string) (found bool, id string) {
    circuit := c.circuit

    for i := 0; i != len(blueprint); i++ {
        for j := 0; j != len(blueprint[i]); j++ {
            if len(circuit) < x + i || len(circuit[i]) < y + j {
                continue
            }

            if blueprint[i][j] == "" {
                continue
            }

            if blueprint[i][j] == "." {
                id = circuit[x + i][y + j]
                continue
            }

            if circuit[x + i][y + j] != blueprint[i][j] {
                found, id = false, ""
                return
            }
        }
    }

    found = true
    return
}

func (c *Circuit) getComponentById(id string) components.Component {
    for _, com := range c.Components {
        if com.Id() == id {
            return com
        }
    }
    return nil
}

func (c *Circuit) getComponentByLocation(x int, y int) components.Component {
    for _, com := range c.Components {
        cX, cY, cWidth, cHeight := com.Space()
        if x >= cX && x <= cX + cWidth && y >= cY && y <= cY + cHeight {
            return com
        }
    }
    return nil
}

func contains(str string, slice []string) bool {
    for _, s := range slice {
        if s == str {
            return true
        }
    }
    return false
}

func (c *Circuit) recognizeIOType(x int, y int) (t string, id string, nx int, ny int, direction int) {
    if x > 0 && len(c.circuit) - 1 > x && y < len(c.circuit[0]) && y > 0 {
        for i := 0; i != 4; i++ {
            mappings := IO_CHECK[i];
            com := c.getComponentByLocation(x + mappings["com_x"], y + mappings["com_y"])
            if com == nil {
                continue
            }
            t = c.circuit[x + mappings["type_x"]][y + mappings["type_y"]]
            if contains(t, com.OutputStreams()) || contains(t, com.InputStreams()) {
                id, nx, ny, direction = com.Id(), x + mappings["x"], y + mappings["y"], i
                return
            }
            t = ""
        }
    }
    return
}

func (c *Circuit) assemble() {
    for y := 0; y != len(c.circuit[0]); y++ {
        for x := 0; x != len(c.circuit); x++ {
            if c.circuit[x][y] != "%" {
                continue
            }

            mio, mid, cX, cY, direction := c.recognizeIOType(x, y)
            mcom := c.getComponentById(mid);
            if !(mcom != nil && contains(mio, mcom.OutputStreams())) {
                continue
            }

            c.connect(mcom, mio, cX, cY, direction)
        }
    }
}

func (c *Circuit) connect(mcom components.Component, mio string, x int, y int, direction int) {
    var char string
    switch direction {
    case 0:
        x++
        if len(c.circuit) <= x {
            break
        }

        char = c.circuit[x][y]

        if match, _ := regexp.MatchString("[|a-zA-Z]", char); match {
            c.connect(mcom, mio, x, y, direction)
        }

        if char == "%" {
            fio, fid, _, _, _ := c.recognizeIOType(x, y)

            if fid != "" {
                fcom := c.getComponentById(fid)
                fcom.Connect(mcom.Output(mio), fio)
                log.Println(fid, fio, "is connected to", mcom.Id(), mio)
            }
        }

        if char == "+" {
            c.connect(mcom, mio, x, y, 2)
            c.connect(mcom, mio, x, y, 3)
            c.connect(mcom, mio, x, y, direction)
        }

        if char == "-" && c.circuit[x + 1][y] == "|" {
            c.connect(mcom, mio, x, y, direction)
        }
    case 1:
        x--
        if x < 0 {
            break
        }

        char = c.circuit[x][y]

        if match, _ := regexp.MatchString("[|a-zA-Z]", char); match {
            c.connect(mcom, mio, x, y, direction)
        }

        if char == "%" {
            fio, fid, _, _, _ := c.recognizeIOType(x, y)

            if fid != "" {
                fcom := c.getComponentById(fid)
                fcom.Connect(mcom.Output(mio), fio)
                log.Println(fid, fio, "is connected to", mcom.Id(), mio)
            }
        }

        if char == "+" {
            c.connect(mcom, mio, x, y, 2)
            c.connect(mcom, mio, x, y, 3)
            c.connect(mcom, mio, x, y, direction)
        }

        if char == "-" && c.circuit[x - 1][y] == "|" {
            c.connect(mcom, mio, x, y, direction)
        }
    case 2:
        y++
        if len(c.circuit[0]) <= y {
            break
        }

        char = c.circuit[x][y]

        if match, _ := regexp.MatchString("[-a-zA-Z]", char); match {
            c.connect(mcom, mio, x, y, direction)
        }

        if char == "%" {
            fio, fid, _, _, _ := c.recognizeIOType(x, y)

            if fid != "" {
                fcom := c.getComponentById(fid)
                fcom.Connect(mcom.Output(mio), fio)
                log.Println(fid, fio, "is connected to", mcom.Id(), mio)
            }
        }

        if char == "+" {
            c.connect(mcom, mio, x, y, 0)
            c.connect(mcom, mio, x, y, 1)
            c.connect(mcom, mio, x, y, direction)
        }

        if char == "|" && c.circuit[x][y + 1] == "-" {
            c.connect(mcom, mio, x, y, direction)
        }
    case 3:
        y--
        if y < 0 {
            break
        }

        char = c.circuit[x][y]

        if match, _ := regexp.MatchString("[-a-zA-Z]", char); match {
            c.connect(mcom, mio, x, y, direction)
        }

        if char == "%" {
            fio, fid, _, _, _ := c.recognizeIOType(x, y)

            if fid != "" {
                fcom := c.getComponentById(fid)
                fcom.Connect(mcom.Output(mio), fio)
                log.Println(fid, fio, "is connected to", mcom.Id(), mio)
            }
        }

        if char == "+" {
            c.connect(mcom, mio, x, y, 0)
            c.connect(mcom, mio, x, y, 1)
            c.connect(mcom, mio, x, y, direction)
        }

        if char == "|" && c.circuit[x][y - 1] == "-" {
            c.connect(mcom, mio, x, y, direction)
        }
    }
}

func (c *Circuit) display() {
    stdscr, err := gc.Init()
    if err != nil {
        log.Println(err)
    }
    defer gc.End()

    rows, cols := stdscr.MaxYX()
    height, width := len(c.circuit) + 1, len(c.circuit[0]) + 1
    y, x := (rows - height) / 2, (cols - width) / 2

    var win *gc.Window
    win, err = gc.NewWindow(height, width, y, x)
    if err != nil {
        log.Println(err)
    }
    defer win.Delete()

    win.Timeout(500)

    for i := 0; i != height - 1; i++ {
        for j := 0; j != width - 1; j++ {
            if c.circuit[i][j] == "" {
                continue
            }
            char := gc.Char([]rune(c.circuit[i][j])[0])
            win.MoveAddChar(i, j, char)
        }
    }

    main:
    for {
        for _, com := range c.Components {
            com.Update()
        }

        for _, com := range c.Components {
            cx, cy, _, _ := com.Space()
            for coord, out := range com.Visual() {
                char := gc.Char([]rune(*out)[0])
                win.MoveAddChar(cx + coord.X, cy + coord.Y, char)
            }
        }

        win.NoutRefresh()
        gc.Update()

        switch win.GetChar() {
            case 'q':
                break main
        }
    }
}
