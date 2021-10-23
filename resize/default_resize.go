//go:build !vips
// +build !vips

package resize

func doFixedResize(img []byte, width, height int) ([]byte, error) {
	panic("Implement me!")
}

func doScaledResize(img []byte, ratio float64) ([]byte, error) {
	panic("Implement me!")
}
