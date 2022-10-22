package exec

import (
	"bytes"
	"fmt"
	"io"

	"github.com/asalih/gika/types"
	"github.com/extrame/ole2"
	"github.com/richardlehane/mscfb"
)

type MsiContentHandler struct {
}

func (z *MsiContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr := bytes.NewReader(context.RawBuffer)

	oledata, err := ole2.Open(rdr, "utf-16")
	if err != nil {
		return nil, err
	}

	dirs, err := oledata.ListDir()
	if err != nil {
		return nil, err
	}
	fmt.Println(dirs)

	for _, dir := range dirs {
		fmt.Println(dir.Name())
	}

	msi, err := mscfb.New(rdr)
	if err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	for {
		entry, err := msi.Next()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(entry)
		if err != nil {
			return nil, err
		}

		entries[entry.Name] = buf.Bytes()
	}

	return entries, nil
}
