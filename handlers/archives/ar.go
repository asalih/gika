package archives

import (
	"bytes"
	"io"

	"github.com/asalih/gika/types"
	"github.com/blakesmith/ar"
)

type ARContentHandler struct {
}

func (t *ARContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	tr := ar.NewReader(context.Reader)

	entries := make(types.Entries)

	for {
		header, err := tr.Next()

		// if no more files are found break
		if err == io.EOF {
			break
		}

		// return any other error
		if err != nil {
			return nil, err
		}

		// if the header is nil, just skip it (not sure how this happens)
		if header == nil {
			continue
		}

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(tr)
		if err != nil {
			return nil, err
		}

		entries[header.Name] = buf.Bytes()

	}

	return entries, nil
}
