package asciiart

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"math"
)

const charOptions = " .:-=+*#%@"

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

	imageWidth := decodedImage.Bounds().Size().X
	imageHeight := decodedImage.Bounds().Size().Y
	for x := 0; x < imageWidth; x++ {
		for y := 0; y < imageHeight; y++ {
			grayScalePixel := color.GrayModel.Convert(decodedImage.At(x, y))
			pixelLightness := getLightness(grayScalePixel)
			fmt.Printf("Lightness at %v, %v, is: %v\n", x, y, pixelLightness)
		}
	}
	return imageBytes, nil
}

// Get the lightness of a colour, based on this source:
// https://stackoverflow.com/a/56678483/6308012
func getLightness(c color.Color) float64 {
	// TODO: ofc this is borked because color.RGBA returns the alpha premultiplied values
	// Need to truncate them down to the last byte of the value, or just replace the
	// first byte with 0's? Does it matter if we just divide by 255 and put it in a float64 after?
	r, g, b, _ := c.RGBA()
	linR := sRGBtoLin(float64(r / 255))
	linG := sRGBtoLin(float64(g / 255))
	linB := sRGBtoLin(float64(b / 255))

	return 0.2126*linR + 0.7152*linG + 0.0722*linB
}

func sRGBtoLin(colorChannel float64) float64 {
	if colorChannel <= 0.04045 {
		return colorChannel / 12.92
	} else {
		return math.Pow((colorChannel+0.055)/1.055, 2.4)
	}
}
