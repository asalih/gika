package archives

import (
	"bytes"

	"github.com/asalih/gika/types"
	"github.com/klauspost/compress/zstd"
)

type ZstdContentHandler struct {
}

func (t *ZstdContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr := bytes.NewReader(context.RawBuffer)

	decoder, err := zstd.NewReader(rdr, zstd.WithDecoderConcurrency(0))
	if err != nil {
		return nil, err
	}
	defer decoder.Close()

	entries := make(types.Entries)

	buf := new(bytes.Buffer)
	buf.ReadFrom(decoder)

	entries[context.FullPath] = buf.Bytes()

	return entries, nil
}
