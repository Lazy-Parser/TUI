package page_generator

import "github.com/charmbracelet/bubbles/key"

// here will be shared things like keys

type keyMapSelection struct {
	Up    key.Binding // move list up
	Down  key.Binding // move list down
	Space key.Binding // select elem
	Enter key.Binding // submit selections
	Right key.Binding // add pool manually
}

var keysSelection = keyMapSelection{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("→", "add pool manually"),
	),
	Space: key.NewBinding(
		key.WithKeys("space", " "),
		key.WithHelp("[Space]", "select item"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("[Enter]", "submit"),
	),
}

type keyMapAddPool struct {
	Up    key.Binding // next focuse
	Down  key.Binding // prev focus
	Enter key.Binding // submit
	Left  key.Binding // move to the exchange generation
}

var keysAddPool = keyMapAddPool{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "previous input"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j", "tab"),
		key.WithHelp("↓/j/tab", "next input"),
	),
	Left: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("←", "previous window"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("[Enter]", "submit"),
	),
}

type keyMapGeneration struct {
	Esc key.Binding
}

var keysGeneration = keyMapGeneration{
	Esc: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("[Esc]", "Menu"),
	),
}
