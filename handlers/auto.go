package gika

import (
	"github.com/asalih/gika/handlers/archives"
	"github.com/asalih/gika/handlers/text"
	"github.com/asalih/gika/types"
)

var contentHandlersMap = map[string]types.IContentHandler{
	"application/zip":             &archives.ZipContentHandler{},
	"application/gzip":            &archives.GzipContentHandler{},
	"application/x-tar":           &archives.TarContentHandler{},
	"text/plain":                  &text.TextContentHandler{},
	"application/x-7z-compressed": &archives.UnarrContentHandler{},
	"application/vnd.rar":         &archives.UnarrContentHandler{},
	"application/x-bzip2":         &archives.BzipContentHandler{},
	"application/x-xz":            &archives.XZContentHandler{},
	"application/zstd":            &archives.ZstdContentHandler{},
	"application/x-compress":      &archives.LzwContentHandler{},
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
