package page_generator

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TODO: write a little info about generator here. Also make a title "Generator" with ascii.
// and in general, make title for each page in header with ascii + a little description what current page do.
func makeBold(str string) string {
	return lipgloss.NewStyle().Bold(true).Render(str)
}

type header struct {
}

func NewHeader() tea.Model {
	return &header{}
}

func (h *header) Init() tea.Cmd                           { return nil }
func (h *header) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return h, nil }
func (h *header) View() string {
	return "Header"
}
