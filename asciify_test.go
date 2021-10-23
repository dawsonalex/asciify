package asciify

import (
	"github.com/stretchr/testify/assert"
	"image/color"
	"testing"
)

func TestLightness(t *testing.T) {
	type testCase struct {
		name              string
		c                 color.Color
		expectedLightness float64
	}

	tests := []testCase{
		{
			name:              "Test with Black",
			c:                 color.Black,
			expectedLightness: 0.0,
		},
		{
			name:              "Test with White",
			c:                 color.White,
			expectedLightness: 1.0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			lightness := getLightness(test.c)
			assert.InDelta(t, test.expectedLightness, lightness, 0.01)
		})
	}
}
