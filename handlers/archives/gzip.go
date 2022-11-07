package archives

import (
	"bytes"
	"compress/gzip"

	"github.com/asalih/gika/types"
)

type GzipContentHandler struct {
}

func (t *GzipContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	gzr, err := gzip.NewReader(context.Reader)
	if err != nil {
		return nil, err
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
