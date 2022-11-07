package archives

import (
	"bytes"

	"github.com/asalih/gika/types"
	"github.com/google/go-cabfile/cabfile"
)

type CabContentHandler struct {
}

func (t *CabContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	cab, err := cabfile.New(context.Reader)
	if err != nil {
		return nil, err
	}
	files := cab.FileList()

	entries := make(types.Entries)

	for _, file := range files {
		fileReader, err := cab.Content(file)
		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(fileReader)
		if err != nil {
			return nil, err
		}

		entries[file] = buf.Bytes()
	}

	return entries, nil
}
