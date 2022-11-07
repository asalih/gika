package doc

import (
	"archive/zip"
	"bytes"
	"os"

	"github.com/asalih/gika/types"
)

type PptxContentHandler struct {
}

func (z *PptxContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr, isRdrAt := context.ReaderAt()
	if !isRdrAt {
		return nil, os.ErrInvalid
	}

	archive, err := zip.NewReader(rdr, context.Size)
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
