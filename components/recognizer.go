package components

type Recognizer interface {
    Blueprint() [][]string
    NewComponent(string, int, int, map[string]string) Component
}
