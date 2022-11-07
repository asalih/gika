package exec

import (
	"bytes"

	"github.com/asalih/gika/types"
	"github.com/asalih/go-msi"
)

type MsiContentHandler struct {
}

func (z *MsiContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	msiPackage, err := msi.Open(context.Reader)
	if err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	streams := msiPackage.Streams()

	for {
		streamName := streams.Next()
		if streamName == "" {
			break
		}

		streamReader, err := msiPackage.ReadStream(streamName)
		if err != nil {
			return nil, err
		}

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(streamReader)
		if err != nil {
			return nil, err
		}

		entries[streamName] = buf.Bytes()
	}

	return entries, nil
}
