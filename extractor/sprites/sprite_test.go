package sprite

import "testing"

func TestBuildSpriteSuccess(t *testing.T) {
	spriteByte := [16]byte{
		// channel 1
		0xF8, // 0b11111000
		0xF8, // 0b11111000
		0xF8, // 0b11111000
		0xF8, // 0b11111000
		0xF8, // 0b11111000
		0x00, // 0b00000000
		0x00, // 0b00000000
		0x00, // 0b00000000

		// channel 2
		0x00, // 0b00000000
		0x00, // 0b00000000
		0x00, // 0b00000000
		0x1F, // 0b00011111
		0x1F, // 0b00011111
		0x1F, // 0b00011111
		0x1F, // 0b00011111
		0x1F, // 0b00011111
	}

	expectArray := [8][8]uint{
		[8]uint{1, 1, 1, 1, 1, 0, 0, 0},
		[8]uint{1, 1, 1, 1, 1, 0, 0, 0},
		[8]uint{1, 1, 1, 1, 1, 0, 0, 0},
		[8]uint{1, 1, 1, 3, 3, 2, 2, 2},
		[8]uint{1, 1, 1, 3, 3, 2, 2, 2},
		[8]uint{0, 0, 0, 2, 2, 2, 2, 2},
		[8]uint{0, 0, 0, 2, 2, 2, 2, 2},
		[8]uint{0, 0, 0, 2, 2, 2, 2, 2},
	}

	sprite := buildSprite(spriteByte)
	actualArray := sprite.ToArrayInt()

	if expectArray != actualArray {
		t.Errorf("BuildSprite() is failed. expect: %v / actual: %v", expectArray, actualArray)
	}
}

func TestOverlapChannelSuccess(t *testing.T) {
	chanel1 := [8]byte{
		0xF8, // 0b11111000
		0xF8, // 0b11111000
		0xF8, // 0b11111000
		0xF8, // 0b11111000
		0xF8, // 0b11111000
		0x00, // 0b00000000
		0x00, // 0b00000000
		0x00, // 0b00000000
	}

	chanel2 := [8]byte{
		0x00, // 0b00000000
		0x00, // 0b00000000
		0x00, // 0b00000000
		0x1F, // 0b00011111
		0x1F, // 0b00011111
		0x1F, // 0b00011111
		0x1F, // 0b00011111
		0x1F, // 0b00011111
	}

	expectArray := [8][8]uint{
		[8]uint{1, 1, 1, 1, 1, 0, 0, 0},
		[8]uint{1, 1, 1, 1, 1, 0, 0, 0},
		[8]uint{1, 1, 1, 1, 1, 0, 0, 0},
		[8]uint{1, 1, 1, 3, 3, 2, 2, 2},
		[8]uint{1, 1, 1, 3, 3, 2, 2, 2},
		[8]uint{0, 0, 0, 2, 2, 2, 2, 2},
		[8]uint{0, 0, 0, 2, 2, 2, 2, 2},
		[8]uint{0, 0, 0, 2, 2, 2, 2, 2},
	}

	actualArray := overlapChannel(chanel1, chanel2)

	if expectArray != actualArray {
		t.Errorf("overlapChannel() is failed. expect: %v / actual: %v", expectArray, actualArray)
	}
}

func TestBuildSpritesSuccess(t *testing.T) {
	characterRom := []byte{
		0xF8, 0xF8, 0xF8, 0xF8, 0xF8, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x1F, 0x1F, 0x1F, 0x1F, 0x1F,
		0xF8, 0xF8, 0xF8, 0xF8, 0xF8, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x1F, 0x1F, 0x1F, 0x1F, 0x1F,
	}

	sprites, err := BuildSprites(characterRom)
	if err != nil {
		t.Fatalf("BuildSprites() is failed. return error: %v", err)
	}

	actualLength := len(sprites)
	if 2 != actualLength {
		t.Errorf("BuildSprites() is failed. expect: %v / actual: %v", 2, actualLength)
	}
}
func TestBuildSpritesNilFailed(t *testing.T) {
	_, err := BuildSprites(nil)
	if err == nil {
		t.Errorf("BuildSprites() is succeeded with nil byte ")
	}
}

func TestBuildSpritesIrregularFailed(t *testing.T) {
	notMultipleOf16Bytes := []byte{
		0xF8, 0xF8, 0xF8, 0xF8, 0xF8, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x1F, 0x1F, 0x1F, 0x1F, 0x1F,
		0xF8, 0xF8, 0xF8, 0xF8, 0xF8, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x1F, 0x1F, 0x1F, 0x1F,
	}
	_, err := BuildSprites(notMultipleOf16Bytes)
	if err == nil {
		t.Errorf("BuildSprites() is succeeded with no multiple of 16 byte ")
	}
}
