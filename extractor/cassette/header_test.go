package cassette

import (
	"testing"
)

func TestNewHeaderSuccess(t *testing.T) {
	prgPageSize, chrPageSize := 1, 2
	romBytes := createHeaderBytes(magicNumber, prgPageSize, chrPageSize)

	header, err := newHeader(romBytes)
	if err != nil {
		t.Fatalf("newHeader() is failed. return error: %v", err)
	}

	if prgPageSize != header.PrgPageSize {
		t.Errorf("newHeader() is failed. program page size is wrong. expect: %v / actual: %v", prgPageSize, header.PrgPageSize)
	}

	if chrPageSize != header.ChrPageSize {
		t.Errorf("newHeader() is failed. character page size is wrong. expect: %v / actual: %v", chrPageSize, header.ChrPageSize)
	}
}

func TestNewHeaderWrongMagicNumber(t *testing.T) {
	romBytes := createHeaderBytes("AAA\x1A", 1, 1)
	_, err := newHeader(romBytes)
	if err == nil {
		t.Errorf("newHeader() is succeeded with wrong magic number ")
	}
}

func createHeaderBytes(magicNumber string, prgPageSize, chrPageSize int) [16]byte {
	var romBytes [16]byte
	copy(romBytes[0:4], []byte(magicNumber))
	romBytes[4] = byte(prgPageSize)
	romBytes[5] = byte(chrPageSize)

	return romBytes
}
