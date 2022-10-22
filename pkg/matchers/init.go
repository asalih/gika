package matchers

import (
	"sync"

	"github.com/h2non/filetype"
)

var syncOnce sync.Once

func Register() {
	syncOnce.Do(func() {
		filetype.AddMatcher(MsiType, MsiMatcher)
	})
}
