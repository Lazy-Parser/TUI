package page_generator

import (
	"github.com/Lazy-Parser/TUI/internal/service"
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

func newOption(title string, desc string) option {
	return option{title: title, desc: desc}
}

type selection struct {
	cursor   int
	selected map[int]struct{}
	opts     []option
	keys     keyMapSelection
}

func NewSelection() *selection {
	return &selection{
		opts:     nil,
		selected: make(map[int]struct{}),
		cursor:   0,
		keys:     keysSelection,
	}
}

func (s *selection) Init() tea.Cmd {
	// request for list of exchanges
	return common.CmdHandler(service.RequestExchangesListMsg{})
}

func (s *selection) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// responces
	case service.ResponseExchangesListMsg:
		return s.handleResponseExchanges(msg)

	// keys
	case tea.KeyMsg:

		switch msg.String() {
		case "q", "ctrl+c":
			return s, tea.Quit
		}

		switch {
		case key.Matches(msg, s.keys.Down):
			return s.handleKeyUp()
		case key.Matches(msg, s.keys.Up):
			return s.handleKeyDown()
		case key.Matches(msg, s.keys.Space):
			return s.handleKeySpace()
		case key.Matches(msg, s.keys.Right):
			return s.handleKeyRight()
		case key.Matches(msg, s.keys.Enter):
			return s.handleKeyEnter()

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

// methods
func (s *selection) handleResponseExchanges(msg service.ResponseExchangesListMsg) (tea.Model, tea.Cmd) {
	res := make([]option, 0, len(msg.Exchanges))
	for _, e := range msg.Exchanges {
		res = append(res, newOption(e, ""))
	}
	s.opts = res

	return s, nil
}

func (s *selection) handleKeyUp() (tea.Model, tea.Cmd) {
	if s.cursor == len(s.opts)-1 {
		s.cursor = 0
	} else {
		s.cursor++
	}

	return s, nil
}

func (s *selection) handleKeyDown() (tea.Model, tea.Cmd) {
	if s.cursor == 0 {
		s.cursor = len(s.opts) - 1
	} else {
		s.cursor--
	}

	return s, nil
}

func (s *selection) handleKeySpace() (tea.Model, tea.Cmd) {
	_, ok := s.selected[s.cursor]
	if ok {
		delete(s.selected, s.cursor)
	} else {
		s.selected[s.cursor] = struct{}{}
	}

	return s, nil
}

func (s *selection) handleKeyEnter() (tea.Model, tea.Cmd) {
	var elems []string
	for idx := range s.selected {
		// idx - selected elems
		elems = append(elems, s.opts[idx].title)
	}

	return s, common.CmdHandler(SelectionModeSubmitMsg{Elems: elems})
}

func (s *selection) handleKeyRight() (tea.Model, tea.Cmd) {
	return s, common.CmdHandler(AddPoolModeMsg{})
}