package asciiart

import (
	"bytes"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"math"
)

const charOptions = " .:-=+*#%@"

// FromImageBuffer returns an ascii image as a slice of bytes.
func FromImageBuffer(width, height int, imageBytes []byte) ([]byte, error) {
	resizer := NewVipsResizer()
	resizedImage, err := resizer.Resize(width, height, imageBytes)
	if err != nil {
		return nil, err
	}

	decodedImage, _, err := image.Decode(bytes.NewBuffer(resizedImage))
	if err != nil {
		return nil, err
	}

	art := make([]byte, 0)
	imageWidth := decodedImage.Bounds().Size().X
	imageHeight := decodedImage.Bounds().Size().Y
	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			grayScalePixel := color.GrayModel.Convert(decodedImage.At(x, y))
			pixelLightness := getLightness(grayScalePixel)
			// better way to map between 0.0-1.0 and 0 and 9?
			if pixelLightness >= 1.0 {
				pixelLightness = 0.9
			}
			charIndex := uint8(pixelLightness * 10)
			art = append(art, charOptions[charIndex])
		}
		art = append(art, '\n')
	}
	return art, nil
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
