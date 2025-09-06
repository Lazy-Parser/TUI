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
	r key.Binding
	q key.Binding

	up        key.Binding
	down      key.Binding
	focusItem key.Binding // select item
	backspace key.Binding // delete elem
	i         key.Binding // more info

	showTokens key.Binding
	showPools  key.Binding
}

func (f footer) Init() tea.Cmd                           { return nil }
func (f footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return f, nil }

func (f footer) helpView() string {
	return f.help.FullHelpView([][]key.Binding{
		{
			f.keymap.q,
			f.keymap.r,
		},
		{
			f.keymap.up,
			f.keymap.down,
		},
		{
			f.keymap.focusItem,
			f.keymap.backspace,
		},
		{
			f.keymap.i,
		},
		{
			f.keymap.showTokens,
			f.keymap.showPools,
		},
	})
}
func (f footer) View() string {
	return f.helpView()
}

func NewFooter() tea.Model {
	return &footer{
		help:   help.New(),
		keymap: keys,
	}
}

var keys = keymap{
	r: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "Reload"),
	),
	q: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "Quit"),
	),
	up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "Up"),
	),
	down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "Down"),
	),
	focusItem: key.NewBinding(
		key.WithKeys("space", " ", "enter"),
		key.WithHelp("[space / enter]", "Select item"),
	),
	backspace: key.NewBinding(
		key.WithKeys("backspace"),
		key.WithHelp("backspace/⌫", "Delete item"),
	),
	i: key.NewBinding(
		key.WithKeys("i", "I"),
		key.WithHelp("i", "Info"),
	),
	showTokens: key.NewBinding(
		key.WithKeys("t", "T"),
		key.WithHelp("t", "Tokens"),
	),
	showPools: key.NewBinding(
		key.WithKeys("p", "P"),
		key.WithHelp("p", "Pools"),
	),
}
