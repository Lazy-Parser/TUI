package page_default

import (
	"tui/internal/tui/components/pages"
)

func NewPageDefault() *pages.Page {
	return pages.NewPage(
		newHeader(),
		newBody(),
		newFooter(),
	)
}
