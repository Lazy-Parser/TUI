package page_generator

import (
	"github.com/Lazy-Parser/TUI/internal/tui/pages"
)

func NewPage() *pages.Page {
	return &pages.Page{
		Header: NewHeader(),
		Main:   NewMain(),
		Footer: NewFooter(),
	}
}
