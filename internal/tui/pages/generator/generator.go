package page_generator

import (
	"github.com/Lazy-Parser/Collector/chains"
	"github.com/Lazy-Parser/Collector/config"
	"github.com/Lazy-Parser/TUI/internal/tui/pages"
)

func NewPage(cfg *config.Config, chainsService *chains.Chains) *pages.Page {
	return &pages.Page{
		Header: NewHeader(cfg),
		Main:   NewMain(cfg, chainsService),
		Footer: NewFooter(),
	}
}
