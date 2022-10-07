package archives

import (
	"archive/tar"
	"bytes"
	"io"

	"github.com/asalih/gika/types"
)

type TarContentHandler struct {
}

func (t *TarContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr := bytes.NewReader(context.RawBuffer)

	tr := tar.NewReader(rdr)

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

		switch header.Typeflag {

		// if it's a file create it
		case tar.TypeReg:
			buf := new(bytes.Buffer)
			_, err := buf.ReadFrom(tr)
			if err != nil {
				return nil, err
			}

			entries[header.Name] = buf.Bytes()
		}
	}

	return entries, nil
}
