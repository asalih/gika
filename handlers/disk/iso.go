package disk

import (
	"bytes"
	"os"
	"strings"

	"github.com/asalih/gika/types"
	"github.com/kdomanski/iso9660"
)

type ISOContentHandler struct {
}

func (t *ISOContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr, isRdrAt := context.ReaderAt()
	if !isRdrAt {
		return nil, os.ErrInvalid
	}

	iso, err := iso9660.OpenImage(rdr)
	if err != nil {
		return nil, err
	}

	root, err := iso.RootDir()
	if err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	err = walk(root, true, "", entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func walk(item *iso9660.File, root bool, parent string, entries types.Entries) error {
	var path string
	if !root {
		path = strings.Join([]string{parent, item.Name()}, "/")
	}

	if item.IsDir() {
		childs, err := item.GetChildren()
		if err != nil {
			return err
		}

		for _, child := range childs {
			walk(child, false, path, entries)
		}
	} else {
		buf := new(bytes.Buffer)
		_, err := buf.ReadFrom(item.Reader())
		if err != nil {
			return err
		}

		entries[path] = buf.Bytes()
	}

	return nil
}
