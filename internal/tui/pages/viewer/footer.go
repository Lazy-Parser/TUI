package page_viewer

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type footer struct {
	keymap keymap
	help   help.Model
}

type keymap struct {
	r    key.Binding
	quit key.Binding

	// up         key.Binding
	// down       key.Binding
	// selectItem key.Binding
}

func (f footer) Init() tea.Cmd                           { return nil }
func (f footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return f, nil }

func (f footer) helpView() string {
	return f.help.ShortHelpView([]key.Binding{
		f.keymap.r,
	})
}
func (f footer) View() string {
	return f.helpView()
}

func NewFooter() tea.Model {
	return &footer{
		help: help.New(),
		keymap: keymap{
			r: key.NewBinding(
				key.WithKeys("r"),
				key.WithHelp("r", "reload"),
			),
		},
	}
}
