package archives

import (
	"bytes"

	"github.com/asalih/gika/types"
	"github.com/ulikunitz/xz"
)

type XZContentHandler struct {
}

func (z *XZContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	archive, err := xz.NewReader(context.Reader)
	if err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(archive)
	if err != nil {
		return nil, err
	}

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
