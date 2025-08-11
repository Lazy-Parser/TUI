package page_viewer

import (
	"github.com/Lazy-Parser/TUI/internal/tui/components/pages"
)

func NewPage() *pages.Page {
	return &pages.Page{
		Header: NewHeader(),
		Main:   NewMain(),
		Footer: NewFooter(),
	}
}
