package text

import (
	"github.com/asalih/gika/types"
)

type TextContentHandler struct {
}

func (t *TextContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	entries := make(types.Entries)
	entries[context.FullPath] = context.RawBuffer

	return entries, nil
}
