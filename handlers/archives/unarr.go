package archives

import (
	"bytes"
	"io"

	"github.com/asalih/gika/types"
	"github.com/gen2brain/go-unarr"
)

type UnarrContentHandler struct {
}

func (u *UnarrContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	entries := make(types.Entries)
	rdr := bytes.NewReader(context.RawBuffer)

	archive, err := unarr.NewArchiveFromReader(rdr)
	if err != nil {
		return nil, err
	}
	defer archive.Close()

	for {
		err := archive.Entry()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		data, err := archive.ReadAll()
		if err != nil {
			return nil, err
		}

		entries[archive.Name()] = data
	}

	return entries, nil
}
