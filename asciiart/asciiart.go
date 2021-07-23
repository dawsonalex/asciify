package asciiart

import (
	"bytes"
	"image"
	"image/color"
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
			// Need to calculate which char to use based on the grayscale RGBA value here.
		}
	}
}
