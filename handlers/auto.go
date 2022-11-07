package gika

import (
	"github.com/asalih/gika/handlers/archives"
	"github.com/asalih/gika/handlers/disk"
	"github.com/asalih/gika/handlers/doc"
	"github.com/asalih/gika/handlers/exec"
	"github.com/asalih/gika/handlers/text"
	"github.com/asalih/gika/types"
)

var contentHandlersMap = map[string]types.IContentHandler{
	"text/plain": &text.TextContentHandler{},

	"application/zip":                   &archives.ZipContentHandler{},
	"application/gzip":                  &archives.GzipContentHandler{},
	"application/x-tar":                 &archives.TarContentHandler{},
	"application/x-7z-compressed":       &archives.UnarrContentHandler{},
	"application/vnd.rar":               &archives.UnarrContentHandler{},
	"application/x-bzip2":               &archives.BzipContentHandler{},
	"application/x-xz":                  &archives.XZContentHandler{},
	"application/zstd":                  &archives.ZstdContentHandler{},
	"application/x-compress":            &archives.LzwContentHandler{},
	"application/x-lzip":                &archives.LzipContentHandler{},
	"application/x-unix-archive":        &archives.ARContentHandler{},
	"application/vnd.ms-cab-compressed": &archives.CabContentHandler{},

	"application/pdf": &doc.PDFContentHandler{},
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document":   &doc.DocxContentHandler{},
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         &doc.XlsxContentHandler{},
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": &doc.PptxContentHandler{},

	"application/x-iso9660-image": &disk.ISOContentHandler{},

	"application/x-msi": &exec.MsiContentHandler{},
}

type AutoDetectContentHandler struct {
}

func (a *AutoDetectContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {

	handler, ok := contentHandlersMap[context.ContentType.Value]
	if !ok {
		return nil, types.ErrUnknownContentType
	}

	return handler.HandleContent(context)
}
