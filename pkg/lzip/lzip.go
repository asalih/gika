package lzip

import (
	"errors"
	"io"

	"github.com/asalih/gika/pkg/lzip/lzma"
)

const (
	minDictionarySize = 1 << 12
	maxDictionarySize = 1 << 29
)

func Decode(reader io.ReadSeeker) (io.Reader, error) {
	dSize, err := validateAndReadSize(reader)
	if err != nil {
		return nil, err
	}
	reader.Seek(0, 0)

	return lzma.ReaderConfig{DictCap: int(dSize)}.NewReader(reader)
}

func validateAndReadSize(reader io.Reader) (int, error) {
	size := 6
	header := make([]byte, size)
	c, err := reader.Read(header)
	if err != nil {
		return 0, err
	}

	if c != size {
		return 0, errors.New("invalid header size")
	}

	if header[0] != 'L' || header[1] != 'Z' || header[2] != 'I' || header[3] != 'P' || header[4] != 1 {
		return 0, errors.New("lzip: invalid header")
	}

	basePower := int(header[5] & 0x1F)
	subtractionNumerator := int((header[5] & 0xE0) >> 5)

	dictSize := (1 << basePower) - subtractionNumerator*(1<<(basePower-4))
	if dictSize < minDictionarySize || dictSize > maxDictionarySize {
		return 0, errors.New("lzip: invalid dictionary size")
	}

	return dictSize, nil
}
