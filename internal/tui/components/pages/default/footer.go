package page_default

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type modelFooter struct {
	keymap keymap
	help   help.Model
}

type keymap struct {
	header key.Binding
	footer key.Binding
	quit   key.Binding

	up         key.Binding
	down       key.Binding
	selectItem key.Binding
}

func (m modelFooter) Init() tea.Cmd {
	return nil
}

func (m modelFooter) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m modelFooter) helpView() string {
	return m.help.ShortHelpView([]key.Binding{
		m.keymap.header,
		m.keymap.footer,
		m.keymap.quit,
		m.keymap.up,
		m.keymap.down,
		m.keymap.selectItem,
	})
}

func (m modelFooter) View() string {
	return m.helpView()
}

func newFooter() modelFooter {
	return modelFooter{
		help: help.New(),
		keymap: keymap{
			header: key.NewBinding(
				key.WithKeys("1"),
				key.WithHelp("1", "header"),
			),
			footer: key.NewBinding(
				key.WithKeys("2"),
				key.WithHelp("2", "footer"),
			),
			quit: key.NewBinding(
				key.WithKeys("q", "ctrl+c"),
				key.WithHelp("q", "quit"),
			),
			up: key.NewBinding(
				key.WithKeys("↑"),
				key.WithHelp("↑", "Move up"),
			),
			down: key.NewBinding(
				key.WithKeys("↓"),
				key.WithHelp("↓", "Move down"),
			),
			selectItem: key.NewBinding(
				key.WithKeys("Enter"),
				key.WithHelp("Enter", "Select item"),
			),
		},
	}
}
