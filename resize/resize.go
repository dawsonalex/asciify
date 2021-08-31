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

// Resizer takes an image and resizes it based on a chosen strategy.
type Resizer struct {
	terminal terminal.Terminal
}

type Options struct {
	resizeStrategy Strategy
	Scale          float64
	Width          int
	Height         int
}

func (r Resizer) Resize(img []byte, options *Options) ([]byte, error) {
	switch options.resizeStrategy {
	case FixedSize:
		return doFixedResize(img, options.Width, options.Height)
	case Scale:
		return doScaledResize(img, options.Scale)
	case Terminal:
		if width, height, err := r.getTerminalSize(); err == nil {
			return doFixedResize(img, width, height)
		} else {
			return nil, err
		}
	default:
		return nil, errors.New("resize: not valid strategy selected")
	}
}

func (r Resizer) getTerminalSize() (int, int, error) {
	t := terminal.NewTerminalAccessor()
	if width, height, err := t.ScreenSize(); err == nil {
		return width, height, nil
	} else {
		return 0, 0, err
	}
}

func doFixedResize(img []byte, width, height int) ([]byte, error) {
	// TODO: Add logic for this.
	return nil, nil
}

func doScaledResize(img []byte, ratio float64) ([]byte, error) {
	// TODO: Add logic for this.
	return nil, nil
}
