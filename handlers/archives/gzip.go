package archives

import (
	"bytes"
	"compress/gzip"

	"github.com/asalih/gika/types"
)

type GzipContentHandler struct {
}

func (t *GzipContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr := bytes.NewReader(context.RawBuffer)

	gzr, err := gzip.NewReader(rdr)
	if err != nil {
		panic(err)
	}
	defer gzr.Close()

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(gzr)
	if err != nil {
		return nil, err
	}

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
