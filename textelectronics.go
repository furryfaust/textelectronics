package main

import (
    "os"
    "log"
    "strings"

    "github.com/furryfaust/textelectronics/circuit"
)

func main() {
    if len(os.Args) < 2 {
        log.Fatal("You need to specify a txt file.")
    }

    values := map[string]string {}
    if len(os.Args) > 2 {
        for i := 2; i != len(os.Args); i++ {
            if pair := strings.Split(os.Args[i], ":"); len(pair) > 1 {
                values[pair[0]] = pair[1]
            }
        }
    }

    options := circuit.NewOptions(os.Args[1]).WithValues(values).WithClassicRecognizers()
    c := circuit.NewCircuit(options)
    c.Run()
}
