package component

type Recognizer interface {
    Blueprint() [][]string
    NewComponent(string, int, int, int, int) Component
}