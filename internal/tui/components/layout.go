package components

import (
	page_default "github.com/Lazy-Parser/TUI/internal/tui/components/pages/default"
	page_viewer "github.com/Lazy-Parser/TUI/internal/tui/components/pages/viewer"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	skyBlue = lipgloss.Color("#4A90E2")
	peach   = lipgloss.Color("#fff")
)

type model struct {
	width, height int

	showHeader   bool
	showFooter   bool
	heightOffset int // because center part of the layout has a border, which adds an extra line
	widthOffset  int

	pageService *PageService
}

func (m model) Init() tea.Cmd {
	// select the first page (which is the default)
	return m.pageService.Init(0)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q":
			return m, tea.Quit

		case "1":
			if m.showHeader {
				m.showHeader = false
				m.heightOffset -= 1
			} else {
				m.showHeader = true
				m.heightOffset += 1
			}
		case "2":
			if m.showFooter {
				m.showFooter = false
				m.heightOffset -= 1
			} else {
				m.showFooter = true
				m.heightOffset += 1
			}
		}

		switch msg.Type {

		case tea.KeyCtrlC:
			return m, tea.Quit

		// go back to the default page
		case tea.KeyEsc:
			m.pageService.SetCurrentPage(0)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width - 2 // offset for border (border goes to the left)
		m.height = msg.Height

	case page_default.PageSelected:
		m.pageService.SetCurrentPage(1)
	}

	// after layout updates, update selected page
	cmd := m.pageService.Update(m.pageService.currentPage, msg)

	if m.pageService.currentPage != 0 {
		// and update default page, because it contains timer
		oneMoreCmd := m.pageService.Update(0, msg)
		cmd = tea.Batch(cmd, oneMoreCmd)
	}

	return m, cmd
}

func (m model) View() string {
	var header string
	if m.showHeader {
		header = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(skyBlue).
			Background(lipgloss.Color(peach)).
			Width(m.width).
			Render(m.pageService.GetCurrentPage().page.Header.View())
	}

	var footer string
	if m.showFooter {
		footer = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(skyBlue).
			Background(lipgloss.Color(peach)).
			Width(m.width).
			Render(m.pageService.GetCurrentPage().page.Footer.View())
	}

	content := lipgloss.NewStyle().
		Width(m.width).
		// accommodate header and footer
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(skyBlue).
		Background(lipgloss.Color(peach)).
		Height(m.height-lipgloss.Height(header)-lipgloss.Height(footer)-m.heightOffset).
		Align(lipgloss.Center, lipgloss.Center).
		Render(m.pageService.GetCurrentPage().page.Main.View())

	return joinComponents(header, content, footer, m)
}

func joinComponents(header, content, footer string, model model) string {
	if !model.showFooter {
		return lipgloss.JoinVertical(lipgloss.Top, header, content)
	}
	if !model.showHeader {
		return lipgloss.JoinVertical(lipgloss.Top, content, footer)
	}
	if !model.showHeader && !model.showFooter {
		return lipgloss.JoinVertical(lipgloss.Top, content)
	}

	return lipgloss.JoinVertical(lipgloss.Top, header, content, footer)
}

func InitLayout() tea.Model {
	payload := []*PageOption{
		NewPageOption(page_default.NewPageDefault()),
		NewPageOption(page_viewer.NewPage()),
	}

	return model{
		pageService:  NewPageService(payload),
		showHeader:   true,
		showFooter:   true,
		heightOffset: 2,
	}
}
