package archives

import (
	"bytes"
	"compress/bzip2"

	"github.com/asalih/gika/types"
)

type BzipContentHandler struct {
}

func (t *BzipContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr := bytes.NewReader(context.RawBuffer)

	bz := bzip2.NewReader(rdr)

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	_, err := buf.ReadFrom(bz)
	if err != nil {
		return nil, err
	}

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
