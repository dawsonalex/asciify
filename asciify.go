package asciify

import (
	"asciify/resize"
	"bytes"
	"github.com/davidbyttow/govips/v2/vips"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

type PixelMapper func(c color.Color, x, y int) byte

type Renderer struct {
	resize   resize.Resizer
	mapPixel PixelMapper
}

func (r Renderer) RenderFromBuffer(buf []byte, options *resize.Options) ([]byte, error) {
	originImage, _, err := image.Decode(bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	width, height, err := r.resolveSize(originImage.Bounds(), options)
	if err != nil {
		return nil, err
	}

	resizedImage, err := r.resize(width, height, buf)
	if err != nil {
		return nil, err
	}

	scaledImage, _, err := image.Decode(bytes.NewBuffer(resizedImage))
	if err != nil {
		return nil, err
	}
	art := make([]byte, 0)
	imageWidth := scaledImage.Bounds().Size().X
	imageHeight := scaledImage.Bounds().Size().Y
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			art = append(art, r.mapPixel(scaledImage.At(x, y), x, y))
		}
		art = append(art, '\n')
	}
	return art, nil
}

func NewRenderer() *Renderer {
	return &Renderer{
		resize:   DefaultResizer(),
		mapPixel: DefaultPixelMapper(),
	}
}

func DefaultResizer() resize.Resizer {
	return func(maxWidth, maxHeight int, imageBuffer []byte) ([]byte, error) {
		vips.Startup(nil)
		defer vips.Shutdown()

		vipsImage, err := vips.NewImageFromBuffer(imageBuffer)
		if err != nil {
			return nil, err
		}

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
}

// TOOD: migrate below logic to this function
func DefaultPixelMapper() PixelMapper {
	return func(c color.Color, x, y int) byte {
		const charOptions = " .:-=+*#%@"

		grayScalePixel := color.GrayModel.Convert(c)
		pixelLightness := getLightness(grayScalePixel)
		// better way to map between 0.0-1.0 and 0 and 9?
		if pixelLightness >= 1.0 {
			pixelLightness = 0.9
		}
		charIndex := uint8(pixelLightness * 10)
		return charOptions[charIndex]
	}
}

// Get the lightness of a colour, based on this source:
// https://stackoverflow.com/a/56678483/6308012
func getLightness(c color.Color) float64 {
	// TODO: this feels gross, make sure I'm not doing something silly here.
	r, g, b, _ := c.RGBA()
	linR := sRGBtoLin(float64(float32(uint8(r)) / 255))
	linG := sRGBtoLin(float64(float32(uint8(g)) / 255))
	linB := sRGBtoLin(float64(float32(uint8(b)) / 255))

	return 0.2126*linR + 0.7152*linG + 0.0722*linB
}

func sRGBtoLin(colorChannel float64) float64 {
	if colorChannel <= 0.04045 {
		return colorChannel / 12.92
	} else {
		return math.Pow((colorChannel+0.055)/1.055, 2.4)
	}
}
