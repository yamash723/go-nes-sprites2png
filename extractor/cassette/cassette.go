package cassette

import "github.com/pkg/errors"

const headerSize int = 0x0010  // 16 byte
const prgUnitSize int = 0x4000 // 16384 byte
const chrUnitSize int = 0x2000 // 8192 byte

// Cassette has header and character rom data
type Cassette struct {
	chrRom []byte
}

// NewCassette is build a cassette struct by byte information of a NES rom file.
func NewCassette(romBytes []byte) (*Cassette, error) {
	if romBytes == nil || len(romBytes) < 16 {
		return nil, errors.New("invalid .nes file. rom file byte length is not enough")
	}

	var headerBytes [16]byte
	copy(headerBytes[:], romBytes[0:16])
	header, err := newHeader(headerBytes)

	if err != nil {
		return nil, errors.Wrap(err, "invalid .nes file. header bytes is wrong")
	}

	startIdx := headerSize + (header.PrgPageSize * prgUnitSize)
	endIdx := startIdx + (header.ChrPageSize * chrUnitSize)
	cassette := &Cassette{chrRom: romBytes[startIdx:endIdx]}

	return cassette, nil
}
