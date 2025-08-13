package components

import (
	"github.com/Lazy-Parser/TUI/internal/tui/components/pages"

	tea "github.com/charmbracelet/bubbletea"
)

type PageOption struct {
	id     int
	page   *pages.Page
	inited bool
}

func (po *PageOption) Init() tea.Cmd {
	if !po.inited {
		po.inited = true
		return po.page.Init()
	}
	return nil
}

func (po *PageOption) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return po.page.Update(msg)
}

func (po *PageOption) SetPage(page *pages.Page) {
	po.page = page
}

func NewPageOption(page *pages.Page) *PageOption {
	return &PageOption{
		inited: false,
		page:   page,
	}
}

// ----
type PageService struct {
	pages       []*PageOption
	currentPage int
}

func NewPageService(pages []*PageOption) *PageService {
	// counter := 0
	return &PageService{
		pages: pages,
	}
}

// Init all components in page with 'i' index and return command
func (ps *PageService) Init(i int) tea.Cmd {
	return ps.pages[i].Init()
}

// Update all components in page with 'i' index and return command
func (ps *PageService) Update(i int, msg tea.Msg) tea.Cmd {
	updatedPage, cmd := ps.pages[i].Update(msg)
	ps.pages[i].SetPage(updatedPage.(*pages.Page))
	return cmd
}

// Update all pages 
func (ps *PageService) UpdateAll(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(ps.pages))
	for i := range ps.pages {
		cmds[i] = ps.Update(i, msg)
	}
	return tea.Batch(cmds...)
}

// Set current page by provided index and init it, so return a command
func (ps *PageService) SetCurrentPage(i int) tea.Cmd {
	ps.currentPage = i
	return ps.pages[i].Init()
}

func (ps *PageService) GetCurrentPage() *PageOption {
	return ps.pages[ps.currentPage]
}

// Return current page index
func (ps *PageService) Index() int {
	return ps.currentPage
}
