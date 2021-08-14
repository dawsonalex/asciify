package asciify

import (
	"bytes"
	"testing"
)

func TestLightness(t *testing.T) {
	input := []byte{0, 255, 0, 255}

	output, err := FromImageBuffer(2, 2, input)
	if err != nil {
		t.Error(err)
		return
	}

	expected := []byte{charOptions[0], charOptions[9], charOptions[0], charOptions[9]}

	if res := bytes.Compare(output, expected); res != 0 {
		t.Errorf("expected: %v, got: %v", expected, output)
		return
	}
}
