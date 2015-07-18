package component

type Component interface {
    Id() string
    Space() (int, int, int, int)
    Update()
    Print()
    Connect(*int, string)
    Output(string) *int
    InputStreams() []string
    OutputStreams() []string
}