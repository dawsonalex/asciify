package asciify

import (
	"github.com/davidbyttow/govips/v2/vips"
)

type Resizer interface {
	Resize(maxWidth, maxHeight int, imageBuffer []byte) ([]byte, error)
}

type vipsResizer struct{}

func NewVipsResizer() *vipsResizer {
	return &vipsResizer{}
}

func (resizer vipsResizer) Resize(maxWidth, maxHeight int, imageBuffer []byte) ([]byte, error) {
	vips.Startup(nil)
	defer vips.Shutdown()

	vipsImage, err := vips.NewImageFromBuffer(imageBuffer)
	if err != nil {
		return nil, err
	}
	// TODO: use maxWidth and maxHeight to calculate the scale for the image.

	hScale, vScale := 1.0, 1.0
	if vipsImage.Width() > maxWidth {
		hScale = float64(maxWidth) / float64(vipsImage.Width())
	}

	if vipsImage.Height() > maxHeight {
		vScale = float64(maxHeight) / float64(vipsImage.Height())
	}
	err = vipsImage.ResizeWithVScale(hScale, vScale, vips.KernelCubic)

	resizedBytes, _, err := vipsImage.ExportJpeg(vips.NewJpegExportParams())
	if err != nil {
		return nil, err
	}

	return resizedBytes, nil
}
