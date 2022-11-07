package archives

import (
	"bytes"

	"github.com/asalih/gika/types"
	"github.com/klauspost/compress/zstd"
)

type ZstdContentHandler struct {
}

func (t *ZstdContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	decoder, err := zstd.NewReader(context.Reader, zstd.WithDecoderConcurrency(0))
	if err != nil {
		return nil, err
	}
	defer decoder.Close()

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(decoder)
	if err != nil {
		return nil, err
	}

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
