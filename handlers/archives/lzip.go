package archives

import (
	"bytes"

	"github.com/asalih/gika/pkg/lzip"
	"github.com/asalih/gika/types"
)

type LzipContentHandler struct {
}

func (z *LzipContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	//TODO: lzip not working, need to fix

	lzr, err := lzip.Decode(context.Reader)
	if err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(lzr)
	if err != nil {
		return nil, err
	}

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
