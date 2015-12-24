package circuit

import (
    "github.com/furryfaust/textelectronics/components"
)

type Options struct {
    Path string
    Recognizers []components.Recognizer
    Values map[string]string
}

func NewOptions(path string) Options {
    o := Options {}
    o.Path = path
    return o
}

func (o Options) WithClassicRecognizers() Options {
    recognizers := []components.Recognizer {
        components.NewAnd2Recognizer(),
        components.NewInputRecognizer(),
        components.NewProbeRecognizer(),
        components.NewOr2Recognizer(),
        components.NewNotRecognizer(),
        components.NewXor2Recognizer(),
        components.NewClockRecognizer(),
        components.NewFlipFlopRecognizer(),
        components.NewHexaconvRecognizer(),
    }
    return o.WithRecognizers(recognizers)
}

func (o Options) WithRecognizers(recognizers []components.Recognizer) Options {
    o.Recognizers = append(o.Recognizers, recognizers...)
    return o
}

func (o Options) WithValues(values map[string]string) Options {
    o.Values = values
    return o
}
