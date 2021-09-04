package resize

import (
	"errors"
	"github.com/qeesung/image2ascii/terminal"
)

type Strategy int

const (
	FixedSize Strategy = iota
	Scale
	Terminal
)

type Options struct {
	resizeStrategy Strategy
	scale          float64
	width          int
	height         int
}

type Option func(options *Options)

func ToScale(scale float64) Option {
	return func(options *Options) {
		options.resizeStrategy = Scale
		options.scale = scale
	}
}

func ToFixed(width, height int) Option {
	return func(options *Options) {
		options.resizeStrategy = FixedSize
		options.width = width
		options.height = height
	}
}

func ToTerminal() Option {
	return func(options *Options) {
		options.resizeStrategy = Terminal
	}
}

// Buffer resizes an image in a byte buffer using the options provided.
func Buffer(img []byte, optionSetter Option) ([]byte, error) {
	options := &Options{
		resizeStrategy: Terminal,
	}
	optionSetter(options)

	switch options.resizeStrategy {
	case FixedSize:
		return doFixedResize(img, options.width, options.height)
	case Scale:
		return doScaledResize(img, options.scale)
	case Terminal:
		if width, height, err := getTerminalSize(); err == nil {
			return doFixedResize(img, width, height)
		} else {
			return nil, err
		}
	default:
		return nil, errors.New("resize: not valid strategy selected")
	}
}

func getTerminalSize() (int, int, error) {
	t := terminal.NewTerminalAccessor()
	if width, height, err := t.ScreenSize(); err == nil {
		return width, height, nil
	} else {
		return 0, 0, err
	}
}
