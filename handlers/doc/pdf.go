package doc

import (
	"bytes"
	"path/filepath"
	"strconv"

	"github.com/asalih/gika/types"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type PDFContentHandler struct {
}

func (t *PDFContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	rdr := bytes.NewReader(context.RawBuffer)

	pdfContext, err := api.ReadContext(rdr, nil)
	if err != nil {
		return nil, err
	}

	if err := api.ValidateContext(pdfContext); err != nil {
		return nil, err
	}

	if err := api.OptimizeContext(pdfContext); err != nil {
		return nil, err
	}

	entries := make(types.Entries)

	metaData, err := pdfContext.ExtractMetadata()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(metaData); i++ {
		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(metaData[i])
		if err != nil {
			return nil, err
		}

		entries[filepath.Join(context.FullPath, "meta_"+strconv.Itoa(i))] = buf.Bytes()
	}

	for i := 1; i <= pdfContext.PageCount; i++ {
		content, err := pdfContext.ExtractPageContent(i)
		if err != nil {
			return nil, err
		}

		name := filepath.Join(context.FullPath, "content_"+strconv.Itoa(i))

		buf := new(bytes.Buffer)
		_, err = buf.ReadFrom(content)
		if err != nil {
			return nil, err
		}

		entries[name] = buf.Bytes()

		imgs, err := pdfContext.ExtractPageImages(i, false)
		if err != nil {
			return nil, err
		}

		for j := 0; j < len(imgs); j++ {
			buf := new(bytes.Buffer)
			_, err = buf.ReadFrom(imgs[j])
			if err != nil {
				return nil, err
			}

			entries[filepath.Join(name, "image_"+strconv.Itoa(j))] = buf.Bytes()
		}
	}

	return entries, nil
}
