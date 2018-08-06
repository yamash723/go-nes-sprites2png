package sprite

// Sprite is type of the NES rom sprite. Sprite has a array of sprite number
type Sprite [8][8]uint

func (s *Sprite) toArrayInt() [8][8]uint {
	return [8][8]uint(*s)
}

// BuildSprite is build sprite type from 16 length bytes
func BuildSprite(spriteBytes [16]byte) *Sprite {
	var channel1, channel2 [8]byte
	copy(channel1[:], spriteBytes[0:8])
	copy(channel2[:], spriteBytes[8:16])

	overlapped := overlapChannel(channel1, channel2)
	sprite := Sprite(overlapped)

	return &sprite
}

func overlapChannel(channel1, channel2 [8]byte) [8][8]uint {
	var result [8][8]uint
	mask := uint(1)

	for y := 0; y < 8; y++ {
		ch1Row := uint(channel1[y]) // 0b11111000
		ch2Row := uint(channel2[y]) // 0b00011111

		for x := uint(0); x < 8; x++ {
			ch1Int := ch1Row >> x & mask
			ch2Int := ch2Row >> x & mask * 2

			xPos := 7 - x
			result[y][xPos] = ch1Int + ch2Int
		}
	}

	return result
}
