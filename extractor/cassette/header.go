package cassette

import (
	"bytes"
	"errors"
)

const magicNumber string = "NES\x1A"

// INesHeader is header of a rom
type INesHeader struct {
	PrgPageSize int
	ChrPageSize int
}

func newHeader(headerBytes [16]byte) (*INesHeader, error) {
	if bytes.Compare(headerBytes[0:4], []byte(magicNumber)) != 0 {
		return nil, errors.New("invalid .nes file. header magic number is wrong")
	}

	herder := &INesHeader{
		PrgPageSize: int(headerBytes[4]),
		ChrPageSize: int(headerBytes[5]),
	}

	return herder, nil
}
