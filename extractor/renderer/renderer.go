package renderer

import (
	"image"
	"image/color"
	"image/draw"

	sprite "github.com/yamash723/go-nes-sprites2png/extractor/sprites"
)

var colorMap = func() map[uint]uint8 {
	return map[uint]uint8{
		0: 0,
		1: 117,
		2: 188,
		3: 255,
	}
}

func colorNumberToPixel(colorNumber uint) color.NRGBA {
	colorCode := colorMap()[colorNumber]
	return color.NRGBA{
		A: 255,
		R: uint8(colorCode),
		G: uint8(colorCode),
		B: uint8(colorCode),
	}
}

// Render is rendering NRGBA image by a sprites
func Render(sprites []*sprite.Sprite, width, height int) *image.NRGBA {
	spritePerLine := width / 8

	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	// all pixel fill with black
	black := color.RGBA{0, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{black}, image.ZP, draw.Src)

	for i, sprite := range sprites {
		spriteStartX := i % spritePerLine * 8
		spriteStartY := i / spritePerLine * 8

		renderSprite(sprite, img, spriteStartX, spriteStartY)
	}

	return img
}

func renderSprite(sprite *sprite.Sprite, img *image.NRGBA, startX, startY int) {
	spriteArray := sprite.ToArrayInt()

	for y, row := range spriteArray {
		for x, colorNumber := range row {
			pixelPosX := startX + x
			pixelPosY := startY + y
			pixel := colorNumberToPixel(colorNumber)

			img.Set(pixelPosX, pixelPosY, pixel)
		}
	}
}
