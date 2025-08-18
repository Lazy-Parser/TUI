package pages

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Component represents a component in the UI like in React
type Page struct {
	Header tea.Model
	Main   tea.Model
	Footer tea.Model
}

// Init initializes all models.
func (p *Page) Init() tea.Cmd {
	cmds := []tea.Cmd{
		p.Header.Init(),
		p.Main.Init(),
		p.Footer.Init(),
	}
	return tea.Batch(cmds...)
}

// Update updates all models based on the given message.
func (p *Page) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	p.Header, cmd = p.Header.Update(msg)
	cmds := []tea.Cmd{cmd}

	p.Main, cmd = p.Main.Update(msg)
	cmds = append(cmds, cmd)

	p.Footer, cmd = p.Footer.Update(msg)
	cmds = append(cmds, cmd)

	return p, tea.Batch(cmds...)
}

func (p *Page) View() string {
	return lipgloss.JoinVertical(lipgloss.Top,
		p.Header.View(),
		p.Main.View(),
		p.Footer.View(),
	)
}

func NewPage(header, main, footer tea.Model) *Page {
	return &Page{Header: header, Main: main, Footer: footer}
}
