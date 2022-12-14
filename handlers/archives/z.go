package archives

import (
	"bytes"

	"github.com/asalih/gika/types"
	"github.com/chuchiy/dcompress"
)

type LzwContentHandler struct {
}

func (t *LzwContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	comp, err := dcompress.NewReader(context.Reader)
	if err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(comp)
	if err != nil {
		return nil, err
	}

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
