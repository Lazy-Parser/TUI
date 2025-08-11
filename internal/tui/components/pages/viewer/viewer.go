package page_viewer

import (
	"tui/internal/tui/components/pages"
)

func NewPage() *pages.Page {
	return &pages.Page{
		Header: NewHeader(),
		Main:   NewMain(),
		Footer: NewFooter(),
	}
}
