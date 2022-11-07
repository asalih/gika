package types

import "errors"

var ErrUnknownContentType = errors.New("unknown content type")
var ErrContentHandlerNotSet = errors.New("content handler not set")
var ErrContextNotSet = errors.New("context not set")

type Entries map[string][]byte

type IContentHandler interface {
	HandleContent(context *GikaContext) (Entries, error)
}

type ContentType struct {
	Value     string
	Extension string
	Charset   string
}
