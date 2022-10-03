package archives

import (
	"bytes"

	"github.com/asalih/gika/types"
	"github.com/chuchiy/dcompress"
)

type LzwContentHandler struct {
}

func (t *LzwContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr := bytes.NewReader(context.RawBuffer)

	comp, err := dcompress.NewReader(rdr)
	if err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	buf.ReadFrom(comp)

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
