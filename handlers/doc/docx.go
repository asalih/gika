package doc

import (
	"archive/zip"
	"bytes"

	"github.com/asalih/gika/types"
)

type DocxContentHandler struct {
}

func (z *DocxContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr := bytes.NewReader(context.RawBuffer)

	archive, err := zip.NewReader(rdr, int64(rdr.Len()))
	if err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	for _, zipFile := range archive.File {
		file, err := zipFile.Open()
		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(file)
		if err != nil {
			return nil, err
		}

		entries[zipFile.Name] = buf.Bytes()

		file.Close()
	}

	return entries, nil
}
