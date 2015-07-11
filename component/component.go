package component

type Component interface {
    Id() string
    Space() (int, int, int, int)
    Update()
    Print()
    Input(string) *int
    Output(string) *int
}