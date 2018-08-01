package cassette

import "testing"

func TestNewCassetteSuccess(t *testing.T) {
	romBytes := createRomBytes(magicNumber, 5, 2)
	cassette, err := NewCassette(romBytes)
	if err != nil {
		t.Fatalf("NewCassette() is failed. return error: %v", err)
	}

	chrRomSize := 2 * chrUnitSize
	if chrRomSize != len(cassette.chrRom) {
		t.Errorf("NewCassette() is failed. character data size is wrong. expect: %v / actual: %v", chrRomSize, cassette.chrRom)
	}
}

func TestNewCassetteNilBytes(t *testing.T) {
	_, err := NewCassette(nil)
	if err == nil {
		t.Errorf("NewCassette() is succeeded with nil byte ")
	}
}

func TestNewCassetteWrongMagicNumber(t *testing.T) {
	romBytes := createRomBytes("AAA\x1A", 1, 1)
	_, err := NewCassette(romBytes)
	if err == nil {
		t.Errorf("NewCassette() is succeeded with wrong magic number ")
	}
}

func TestNewCassetteNotEnoughLengthByte(t *testing.T) {
	romBytes := make([]byte, 15)
	copy(romBytes[0:4], magicNumber)

	_, err := NewCassette(romBytes)
	if err == nil {
		t.Errorf("NewCassette() is succeeded with not enough length byte ")
	}
}

func createRomBytes(magicNumber string, prgPageSize, chrPageSize int) []byte {
	prgDataSize := prgPageSize * prgUnitSize
	chrDataSize := chrPageSize * chrUnitSize
	romSize := headerSize + prgDataSize + chrDataSize

	romBytes := make([]byte, romSize)
	copy(romBytes[0:4], []byte(magicNumber))
	romBytes[4] = byte(prgPageSize)
	romBytes[5] = byte(chrPageSize)

	return romBytes
}
