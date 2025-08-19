package tui

import (
	"github.com/Lazy-Parser/Collector/chains"
	"github.com/Lazy-Parser/Collector/config"
	"github.com/Lazy-Parser/Collector/market"
	"github.com/Lazy-Parser/TUI/internal/tui/command"
	"github.com/Lazy-Parser/TUI/internal/tui/pages"
	page_default "github.com/Lazy-Parser/TUI/internal/tui/pages/default"
	page_generator "github.com/Lazy-Parser/TUI/internal/tui/pages/generator"
	page_viewer "github.com/Lazy-Parser/TUI/internal/tui/pages/viewer"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO
// internal/tui/
// ├── managers/           # Resource and state managers
// │   ├── resource.go
// │   └── page.go
// ├── handlers/           # Specialized message handlers
// │   ├── common.go
// │   └── navigation.go
// ├── events/             # Event system
// │   └── bus.go
// ├── components/         # Reusable UI components
// └── pages/              # Page implementations
//
// Architecture Recommendations
// Use Delegation: Let your main model delegate to specialized handlers
// Implement Interfaces: Allow components to opt into behaviors
// Centralize State: Create a single source of truth for application data
// Add Event System: Use pub/sub for loose coupling between components
// Error Boundaries: Handle errors at appropriate levels

var (
	skyBlue = lipgloss.Color("#4A90E2")
	// peach   = lipgloss.Color("#fff")
)

type model struct {
	width, height int

	showHeader   bool
	showFooter   bool
	heightOffset int // because center part of the layout has a border, which adds an extra line
	widthOffset  int

	pageService *pages.PageService
}

func (m model) Init() tea.Cmd {
	// select the first page (which is the default)
	return m.pageService.Init(0)
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q":
			return m, command.OnQuit()

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
			cmd := command.OnQuit()
			return m, cmd

		// go back to the default page
		case tea.KeyEsc:
			m.pageService.SetCurrentPage(0)
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width - 2 // offset for border (border goes to the left)
		m.height = msg.Height
		// window size msg is very important, so update all pages
		cmd := m.pageService.UpdateAll(msg)
		return m, cmd

	case page_default.PageSelected:
		switch msg.Title {
		case "Database":
			cmd := m.pageService.SetCurrentPage(1)
			return m, cmd

		case "Generator":
			cmd := m.pageService.SetCurrentPage(2)
			return m, cmd
		}

	}

	// after layout updates, update selected page
	cmd := m.pageService.Update(m.pageService.CurrentPageIdx(), msg)

	if m.pageService.CurrentPageIdx() != 0 {
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
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(skyBlue).
			Width(m.width).
			Render(m.pageService.GetCurrentPage().Page.Header.View())
	}

	var footer string
	if m.showFooter {
		footer = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Border(lipgloss.NormalBorder(), true).
			BorderForeground(skyBlue).
			Width(m.width).
			Render(m.pageService.GetCurrentPage().Page.Footer.View())
	}

	content := lipgloss.NewStyle().
		Width(m.width).
		// accommodate header and footer
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(skyBlue).
		Height(m.height-lipgloss.Height(header)-lipgloss.Height(footer)-m.heightOffset).
		Align(lipgloss.Center, lipgloss.Center).
		Render(m.pageService.GetCurrentPage().Page.Main.View())

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

func InitLayout(tokenRepo market.TokenRepo, cfg *config.Config, chainsService *chains.Chains) tea.Model {
	payload := []*pages.PageOption{
		pages.NewPageOption(page_default.NewPageDefault()),
		pages.NewPageOption(page_viewer.NewPage(tokenRepo)),
		pages.NewPageOption(page_generator.NewPage(cfg, chainsService)),
	}

	return &model{
		pageService:  pages.NewPageService(payload),
		showHeader:   true,
		showFooter:   true,
		heightOffset: 2,
	}
}
