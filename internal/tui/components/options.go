package component

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type keyMap struct {
	Up    key.Binding // move list up
	Down  key.Binding // move list down
	Space key.Binding // select elem
	Enter key.Binding // select elem
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Space: key.NewBinding(
		key.WithKeys("space", " "),
		key.WithHelp("space", "select item"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "submit"),
	),
}


type OptionsSubmitMsg struct {
	Elems []string // titles from options
}

type option struct {
	title string
	desc string
}

type Options struct {
	cursor int
	selected map[int]struct{}
	opts []option
	keys keyMap
}

func (o *Options) Init() tea.Cmd {
	return nil
}

func (o *Options) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, o.keys.Down):
			if o.cursor == len(o.opts) - 1 { // if already the last elem selected
				o.cursor = 0
			} else {
				o.cursor--
			}
		
		case key.Matches(msg, o.keys.Up):
			if o.cursor == 0 {
				o.cursor = len(o.opts) - 1 // set the last elem
			} else {
				o.cursor++
			}
			
		case key.Matches(msg, o.keys.Space): 
			_, ok := o.selected[o.cursor]
			if ok {
				delete(o.selected, o.cursor)
			} else {
				o.selected[o.cursor] = struct{}{}
			}
			
		case key.Matches(msg, o.keys.Enter):
			var elems []string
			for idx := range o.selected {
				// idx - selected elems
				elems = append(elems, o.opts[idx].title)
			}
			
			msg := OptionsSubmitMsg{
				Elems: elems,
			}
			
			return o, MsgHandler(msg)
		}
	}

	return o, nil
}

func (o *Options) View() string {
	var str string
	
	for i, opt := range o.opts {
		_, isSelected := o.selected[i]
		var cursor string
		if isSelected {
			cursor = "x"
		} else {
			cursor = " "
		}
		
		
		str += "[" + cursor + "] " + opt.title + "\n"
	}
	
	return str
}

// Add option to the list
func NewOption(title string, desc string) option {
	return option{title: title, desc: desc}
}

func NewOptionsFrame(opts ...option) *Options {
	return &Options{
		opts: opts,
		selected: make(map[int]struct{}),
		cursor: 0,
		keys: keys,
	}
}

func MsgHandler(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}