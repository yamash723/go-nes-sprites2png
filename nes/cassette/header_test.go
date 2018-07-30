package cassette_test

import (
	"testing"

	"github.com/yamash723/go-nes-sprites2png/nes/cassette"
)

func TestBuildFromBytesSuccess(t *testing.T) {
	prgPageSize, chrPageSize := 1, 2
	romBytes := createRomBytes("NES\x1A", prgPageSize, chrPageSize)

	header, err := cassette.BuildFromBytes(romBytes)
	if err != nil {
		t.Fatalf("BuildFromBytes() is failed. return error: %v", err)
	}

	if prgPageSize != header.PrgPageSize {
		t.Errorf("BuildFromBytes() is failed. program page size is wrong. expect: %v / actual: %v", prgPageSize, header.PrgPageSize)
	}

	if chrPageSize != header.ChrPageSize {
		t.Errorf("BuildFromBytes() is failed. character page size is wrong. expect: %v / actual: %v", chrPageSize, header.ChrPageSize)
	}
}

func TestBuildFromBytesWrongMagicNumber(t *testing.T) {
	romBytes := createRomBytes("AAA\x1A", 1, 1)
	_, err := cassette.BuildFromBytes(romBytes)
	if err == nil {
		t.Errorf("BuildFromBytes() is succeeded with wrong magic number ")
	}
}

func createRomBytes(magicNumber string, prgPageSize, chrPageSize int) [16]byte {
	var romBytes [16]byte
	copy(romBytes[0:4], []byte(magicNumber))
	romBytes[4] = byte(prgPageSize)
	romBytes[5] = byte(chrPageSize)

	return romBytes
}
