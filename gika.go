package gika

import (
	"io"
	"os"

	"github.com/asalih/gika/pkg/matchers"
	"github.com/asalih/gika/types"
)

type Gika struct {
	handle *Handle
}

type Handle struct {
	gikaContext    *types.GikaContext
	contentHandler types.IContentHandler
}

func New(contentHandler types.IContentHandler) *Gika {
	//register custom matchers
	matchers.Register()

	ctx := types.NewContext()

	return &Gika{
		handle: &Handle{
			gikaContext:    ctx,
			contentHandler: contentHandler,
		},
	}
}

//WithPath sets the path of the file to be processed
func (g *Gika) WithPath(path string) (*Handle, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		return nil, err
	}

	_, err = g.handle.gikaContext.Update(path, stat.Size(), file)
	if err != nil {
		return nil, err
	}

	return g.handle, nil
}

//WithReader sets the reader of the file to be processed
func (g *Gika) WithReader(size int64, reader io.ReadSeekCloser) (*Handle, error) {
	_, err := g.handle.gikaContext.Update("", size, reader)
	if err != nil {
		return nil, err
	}

	return g.handle, nil
}

func (g *Handle) Read() (types.Entries, error) {
	if g.contentHandler == nil {
		return nil, types.ErrContentHandlerNotSet
	}

	if g.gikaContext == nil {
		return nil, types.ErrContextNotSet
	}

	return g.contentHandler.HandleContent(g.gikaContext)
}

func (g *Handle) Close() error {
	return g.gikaContext.Close()
}
