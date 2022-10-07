package archives

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"

	"github.com/asalih/gika/types"
	"github.com/ulikunitz/xz/lzma"
)

type LzipContentHandler struct {
}

func (z *LzipContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	//TODO: lzip not working, need to fix
	rdr := bytes.NewReader(context.RawBuffer[5:])

	size, err := validateAndReadSize(context.RawBuffer)
	if err != nil {
		return nil, err
	}

	props := getProperties(size)

	lsize := binary.LittleEndian.Uint32(props[1:])
	fmt.Println(lsize)
	fmt.Println(size)
	fmt.Println(props)
	fmt.Println(len(props) < 5)

	archive, err := lzma.ReaderConfig{DictCap: size}.NewReader(rdr)
	if err != nil {
		return nil, err
	}
	fmt.Println(archive)

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(archive)
	if err != nil {
		return nil, err
	}

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}

func validateAndReadSize(buf []byte) (int, error) {
	if len(buf) < 5 {
		return 0, errors.New("lzma: invalid header")
	}
	if buf[0] != 'L' || buf[1] != 'Z' || buf[2] != 'I' || buf[3] != 'P' || buf[4] != 1 {
		return 0, errors.New("lzma: invalid header")
	}

	basePower := int(buf[5] & 0x1F)
	subtractionNumerator := int((buf[5] & 0xE0) >> 5)

	return (1 << basePower) - subtractionNumerator*(1<<(basePower-4)), nil
}

func getProperties(dictSize int) []byte {
	return []byte{
		93,
		// Dictionary size as 4-byte little-endian value
		byte(dictSize & 0xff),
		byte((dictSize >> 8) & 0xff),
		byte((dictSize >> 16) & 0xff),
		byte((dictSize >> 24) & 0xff),
	}
}
