package archives

import (
	"bytes"
	"compress/zlib"

	"github.com/asalih/gika/types"
)

type ZlibContentHandler struct {
}

func (t *ZlibContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	archive, err := zlib.NewReader(context.Reader)
	if err != nil {
		return nil, err
	}
	defer archive.Close()

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(archive)
	if err != nil {
		return nil, err
	}

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
