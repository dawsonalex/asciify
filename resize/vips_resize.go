//go:build vips
// +build vips

package resize

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/qeesung/image2ascii/terminal"
)

func doFixedResize(img []byte, width, height int) ([]byte, error) {
	vips.Startup(nil)
	defer vips.Shutdown()

	vipsImage, err := vips.NewImageFromBuffer(img)
	if err != nil {
		return nil, err
	}

	hScale, vScale := 1.0, 1.0
	if vipsImage.Width() > width {
		hScale = float64(width) / float64(vipsImage.Width())
	}

	if vipsImage.Height() > height {
		vScale = float64(height) / float64(vipsImage.Height())
	}
	err = vipsImage.ResizeWithVScale(hScale, vScale, vips.KernelCubic)

	resizedBytes, _, err := vipsImage.ExportJpeg(vips.NewJpegExportParams())
	if err != nil {
		return nil, err
	}

	return resizedBytes, nil
}

func doScaledResize(img []byte, ratio float64) ([]byte, error) {
	vips.Startup(nil)
	defer vips.Shutdown()

	vipsImage, err := vips.NewImageFromBuffer(img)
	if err != nil {
		return nil, err
	}

	err = vipsImage.ResizeWithVScale(ratio, ratio, vips.KernelCubic)

	resizedBytes, _, err := vipsImage.ExportJpeg(vips.NewJpegExportParams())
	if err != nil {
		return nil, err
	}

	return resizedBytes, nil
}

func doResizeToTerminal(img []byte, t terminal.Terminal) ([]byte, error) {
	return nil, nil
}
