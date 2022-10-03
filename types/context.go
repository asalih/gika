package types

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

type GikaContext struct {
	FullPath    string
	RawBuffer   []byte
	ContentType ContentType
}

func NewGikaContext(fullPath string, buf []byte) (*GikaContext, error) {
	var contentType ContentType

	kind, err := filetype.Match(buf)
	if err != nil {
		return nil, err
	}

	if kind != filetype.Unknown {
		contentType = ContentType{
			Extension: kind.Extension,
			Value:     kind.MIME.Value,
		}
	} else {
		//try to get content type from http content detector api
		rawType := http.DetectContentType(buf)
		charset := ""
		if idx := strings.Index(rawType, ";"); idx > 0 {
			charset = strings.Trim(rawType[idx+1:], " ")
			rawType = rawType[:idx]
		}

		contentType = ContentType{
			Extension: strings.TrimLeft(filepath.Ext(fullPath), "."),
			Value:     rawType,
			Charset:   charset,
		}
	}

	if contentType.Value == "" || contentType.Value == "application/octet-stream" {
		return nil, ErrUnknownContentType
	}

	return &GikaContext{
		FullPath:    fullPath,
		RawBuffer:   buf,
		ContentType: contentType,
	}, nil
}
