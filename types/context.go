package types

import (
	"io"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/h2non/filetype"
)

const (
	maxOffset = 32774
)

type GikaContext struct {
	FullPath    string
	Size        int64
	ContentType ContentType

	Reader io.ReadSeekCloser
}

func NewContext() *GikaContext {
	return &GikaContext{}
}

func (c *GikaContext) Update(fullPath string, size int64, reader io.ReadSeekCloser) (*GikaContext, error) {
	var contentType ContentType

	// Seek to the beginning of the file
	_, err := reader.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	// detect content type
	// currently biggest offset belongs to iso9660 file so we need to read 32k bytes
	var header = make([]byte, maxOffset)
	_, err = reader.Read(header)
	if err != nil {
		return nil, err
	}

	// Seek to the beginning of the file
	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		return nil, err
	}

	kind, err := filetype.Match(header)
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
		rawType := http.DetectContentType(header)
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

	c.ContentType = contentType
	c.FullPath = fullPath
	c.Size = size
	c.Reader = reader

	return c, nil
}

func (c *GikaContext) Close() error {
	return c.Reader.Close()
}

func (c *GikaContext) ReaderAt() (io.ReaderAt, bool) {
	r, ok := c.Reader.(io.ReaderAt)
	return r, ok
}
