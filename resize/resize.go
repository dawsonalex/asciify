package resize

import (
	"errors"
	"github.com/qeesung/image2ascii/terminal"
)

type Strategy int

const (
	fixedSizeStrategy Strategy = iota
	scaleStrategy
	terminalStrategy
)

type options struct {
	resizeStrategy Strategy
	scale          float64
	width          int
	height         int
}

type Option func(options *options)

func ToScale(scale float64) Option {
	return func(options *options) {
		options.resizeStrategy = scaleStrategy
		options.scale = scale
	}
}

func ToFixed(width, height int) Option {
	return func(options *options) {
		options.resizeStrategy = fixedSizeStrategy
		options.width = width
		options.height = height
	}
}

func ToTerminal() Option {
	return func(options *options) {
		options.resizeStrategy = terminalStrategy
	}
}

// Buffer resizes an image in a byte buffer using the options provided.
func Buffer(img []byte, optionSetter Option) ([]byte, error) {
	options := &options{
		resizeStrategy: terminalStrategy,
	}
	optionSetter(options)

	switch options.resizeStrategy {
	case fixedSizeStrategy:
		return doFixedResize(img, options.width, options.height)
	case scaleStrategy:
		return doScaledResize(img, options.scale)
	case terminalStrategy:
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
