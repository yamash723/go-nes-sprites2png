package renderer

import (
	"image/color"
	"testing"
)

func TestColorNumberToPixelSuccess(t *testing.T) {

	testExecute := func(colorNumber uint, colorCode uint8) {
		actualColor := colorNumberToPixel(colorNumber)
		expectColor := color.NRGBA{
			A: 255,
			R: uint8(colorCode),
			G: uint8(colorCode),
			B: uint8(colorCode),
		}
		if expectColor != actualColor {
			t.Errorf("colorNumberToPixel() is failed. colorNumber: %v / expect: %v / actual: %v", colorNumber, expectColor, actualColor)
		}
	}

	testExecute(0, 0)   // black
	testExecute(1, 117) // deep gray
	testExecute(2, 188) // light gray
	testExecute(3, 255) // white
}
