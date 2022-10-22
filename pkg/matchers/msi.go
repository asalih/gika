package matchers

import (
	"bytes"

	"github.com/h2non/filetype"
)

var MsiType = filetype.NewType("msi", "application/x-msi")
var msiMagicBytes = []byte{0xD0, 0xCF, 0x11, 0xE0, 0xA1, 0xB1, 0x1A, 0xE1}

func MsiMatcher(buf []byte) bool {
	return len(buf) > 8 && bytes.Equal(buf[0:8], msiMagicBytes)
}
