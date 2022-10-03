package types

import "errors"

var ErrUnknownContentType = errors.New("unknown content type")

type Entries map[string][]byte

type IContentHandler interface {
	HandleContent(context *GikaContext) (Entries, error)
}

type ContentType struct {
	Value     string
	Extension string
	Charset   string
}
