package asciiart

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
	err = vipsImage.ResizeWithVScale(0.1, 0.05, vips.KernelAuto)

	resizedBytes, _, err := vipsImage.ExportJpeg(vips.NewJpegExportParams())
	if err != nil {
		return nil, err
	}

	return resizedBytes, nil
}
