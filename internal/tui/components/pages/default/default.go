package page_default

import (
	"github.com/Lazy-Parser/TUI/internal/tui/components/pages"
)

func NewPageDefault() *pages.Page {
	return pages.NewPage(
		newHeader(),
		newBody(),
		newFooter(),
	)
}
