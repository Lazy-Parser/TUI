package page_viewer

import (
	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/TUI/internal/tui/components/pages"
)

func NewPage(tokenRepo market.TokenRepo) *pages.Page {
	return &pages.Page{
		Header: NewHeader(),
		Main:   NewMain(tokenRepo),
		Footer: NewFooter(),
	}
}
