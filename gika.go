package gika

import (
	"os"

	"github.com/asalih/gika/types"
)

type Gika struct {
	context        *types.GikaContext
	contentHandler types.IContentHandler
}

//New creates new Gika instance with given path
func New(contentHandler types.IContentHandler, path string) (*Gika, error) {
	buf, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	ctx, err := types.NewGikaContext(path, buf)
	if err != nil {
		return nil, err
	}

	return &Gika{
		contentHandler: contentHandler,
		context:        ctx,
	}, nil
}

func NewWithBuffer(contentHandler types.IContentHandler, buf []byte) (*Gika, error) {
	ctx, err := types.NewGikaContext("", buf)
	if err != nil {
		return nil, err
	}

	return &Gika{
		contentHandler: contentHandler,
		context:        ctx,
	}, nil
}

func (g *Gika) Read() (types.Entries, error) {
	return g.contentHandler.HandleContent(g.context)
}
