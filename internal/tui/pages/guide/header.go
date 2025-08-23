package page_guide

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	center = lipgloss.
		NewStyle().
		Align(lipgloss.Center)

	titleStyle = lipgloss.
			NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#00FF00"))

	title = `
  ____ _   _ ___ ____  _____
 / ___| | | |_ _|  _ \| ____|
| |  _| | | || || | | |  _|
| |_| | |_| || || |_| | |___
 \____|\___/|___|____/|_____|
`
)

type header struct{}

func NewHeader() tea.Model { return &header{} }

func (h *header) Init() tea.Cmd                           { return nil }
func (h *header) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return h, nil }

func (h *header) View() string {
	return center.Render(titleStyle.Render(title))
}
