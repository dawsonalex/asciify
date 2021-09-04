package asciify

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

type PixelMapper func(c color.Color, x, y int) byte

func RenderFromBuffer(buf []byte, pixelMapper PixelMapper) ([]byte, error) {
	scaledImage, _, err := image.Decode(bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}
	art := make([]byte, 0)
	imageWidth := scaledImage.Bounds().Size().X
	imageHeight := scaledImage.Bounds().Size().Y
	for y := 0; y < imageHeight; y++ {
		for x := 0; x < imageWidth; x++ {
			art = append(art, pixelMapper(scaledImage.At(x, y), x, y))
		}
		art = append(art, '\n')
	}
	return art, nil
}

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
