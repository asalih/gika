package doc

import (
	"fmt"

	"github.com/asalih/gika/types"
)

type PDFContentHandler struct {
}

func (t *PDFContentHandler) HandleContent(context *types.GikaContext) (types.Entries, error) {
	return nil, fmt.Errorf("not implemented")
}
