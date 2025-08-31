package page_generator

import (
	"log"

	"github.com/Lazy-Parser/TUI/internal/tui/common"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type SelectionModeSubmitMsg struct {
	Elems []string
}

var (
	border = lipgloss.
		NewStyle().
		Border(lipgloss.NormalBorder(), true).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
)

type option struct {
	title string
	desc  string
}

type selection struct {
	cursor   int
	selected map[int]struct{}
	opts     []option
	keys     keyMapSelection
}

func (s *selection) Init() tea.Cmd {
	return nil
}

func (s *selection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		log.Printf("Pressed: %q (Type=%v)", msg.String(), msg.Type)

		switch msg.String() {
		case "q", "ctrl+c":
			return s, tea.Quit
		}

		switch {

		case key.Matches(msg, s.keys.Down):
			if s.cursor == len(s.opts)-1 {
				s.cursor = 0
			} else {
				s.cursor++
			}

		case key.Matches(msg, s.keys.Up):
			if s.cursor == 0 {
				s.cursor = len(s.opts) - 1
			} else {
				s.cursor--
			}

		case key.Matches(msg, s.keys.Space):
			_, ok := s.selected[s.cursor]
			if ok {
				delete(s.selected, s.cursor)
			} else {
				s.selected[s.cursor] = struct{}{}
			}

		case key.Matches(msg, s.keys.Right):
			return s, common.CmdHandler(AddPoolModeMsg{})

		case key.Matches(msg, s.keys.Enter):
			log.Println("Enter selection fire")
			var elems []string
			for idx := range s.selected {
				// idx - selected elems
				elems = append(elems, s.opts[idx].title)
			}

			return s, common.CmdHandler(SelectionModeSubmitMsg{Elems: elems})
		}
	}

	return s, nil
}

func (s *selection) View() string {
	var str string

	str += "Select exchanges:\n\n"
	for i, opt := range s.opts {
		_, isSelected := s.selected[i]
		var cursor string
		if isSelected {
			cursor = "x"
		} else {
			cursor = " "
		}

		elem := "[" + cursor + "] " + opt.title
		if s.cursor == i {
			str += focused(elem)
		} else {
			str += elem
		}
		str += "\n"
	}
	str += "\n"
	str += "Submit [Enter]\n"
	str += "Want to add manually? Press [→]"

	str = border.
		Height(lipgloss.Height(str) + 5). // 6 - padding
		Width(lipgloss.Width(str) + 20).  // 10 - padding
		Render(str)

	return str
}
func focused(str string) string {
	return "-> " + str + " <-"
}

// Add option to the list
func NewOption(title string, desc string) option {
	return option{title: title, desc: desc}
}

func NewSelection(opts ...option) *selection {
	return &selection{
		opts:     opts,
		selected: make(map[int]struct{}),
		cursor:   0,
		keys:     keysSelection,
	}
}
