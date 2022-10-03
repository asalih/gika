package archives

import (
	"archive/zip"
	"bytes"

	"github.com/asalih/gika/types"
)

type ZipContentHandler struct {
}

func (z *ZipContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
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
		buf.ReadFrom(file)
		entries[zipFile.Name] = buf.Bytes()

		file.Close()
	}

	return entries, nil
}
