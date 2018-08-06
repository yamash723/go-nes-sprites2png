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

	sprite := BuildSprite(spriteByte)
	actualArray := sprite.toArrayInt()

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
