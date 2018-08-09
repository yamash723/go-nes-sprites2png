package extractor

import (
	"image/png"
	"io/ioutil"
	"os"

	"github.com/yamash723/go-nes-sprites2png/extractor/cassette"
	"github.com/yamash723/go-nes-sprites2png/extractor/renderer"
	"github.com/yamash723/go-nes-sprites2png/extractor/sprites"

	"github.com/pkg/errors"
)

// Execute is extract png image in which sprite is drawn
func Execute(romPath, outputPath string, width, height int) error {
	romByte, err := ioutil.ReadFile(romPath)
	if err != nil {
		errors.Wrap(err, "read a rom file failed.")
	}

	cassette, err := cassette.NewCassette(romByte)
	if err != nil {
		errors.Wrap(err, "rom file content is wrong.")
	}

	sprites, err := sprite.BuildSprites(cassette.CharacterRom)
	if err != nil {
		errors.Wrap(err, "create sprites failed.")
	}

	img := renderer.Render(sprites, width, height)

	file, err := os.Create(outputPath)
	if err != nil {
		errors.Wrap(err, "can't open outputPath.")
	}
	defer file.Close()

	if err := png.Encode(file, img); err != nil {
		errors.Wrap(err, "can't encode sprite image.")
	}

	return nil
}
