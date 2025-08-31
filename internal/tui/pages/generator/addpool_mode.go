package page_generator

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Lazy-Parser/TUI/internal/tui/common"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type AddPoolModeSubmitMsg struct {
	Address string
	Network string
}
type clearWarningFieldMsg struct{}

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle
	noStyle      = lipgloss.NewStyle()

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))

	warning = lipgloss.NewStyle().Foreground(lipgloss.Color("#ee6622"))
)

type AddPool struct {
	focus   int // 0 - address, 1 - network
	inputs  []textinput.Model
	warning string
	keys    keyMapAddPool
}

func (model *AddPool) Init() tea.Cmd {
	return textinput.Blink
}

func (model *AddPool) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return model, tea.Quit
		}

		switch {
		case key.Matches(msg, model.keys.Down):
			return model.handleDown()
		case key.Matches(msg, model.keys.Up):
			return model.handleUp()
		case key.Matches(msg, model.keys.Enter):
			return model.handleEnter()
		case key.Matches(msg, model.keys.Left):
			return model.handleSetSelectionMode()
		}
	case clearWarningFieldMsg:
		model.warning = ""

	case AddPoolModeSubmitMsg:
		log.Println(msg)
	}

	cmd := model.updateInputs(msg)
	return model, cmd
}

func (model *AddPool) View() string {
	var b strings.Builder

	if model.warning != "" {
		warn := "(" + model.warning + ")"
		b.WriteString(warning.Render(warn))
		b.WriteString("\n")
	}

	b.WriteString("To add new pool manually, fill next inputs:\n\n")
	for i := range model.inputs {
		b.WriteString(model.inputs[i].View())
		if i < len(model.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if model.focus == len(model.inputs)-1 {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString("Want to go back? Press [←]")

	str := border.
		Height(lipgloss.Height(b.String()) + 5). // 5 - padding
		Width(lipgloss.Width(b.String()) + 20).  // 20 - padding
		Render(b.String())
	return str
}

// methods
func (model *AddPool) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(model.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range model.inputs {
		model.inputs[i], cmds[i] = model.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (model *AddPool) handleDown() (*AddPool, tea.Cmd) {
	if model.focus == len(model.inputs)-1 {
		model.focus = 0
	} else {
		model.focus++
	}

	var cmd tea.Cmd
	model, cmd = model.changeFocusUpdate()

	return model, cmd
}

func (model *AddPool) handleUp() (*AddPool, tea.Cmd) {
	if model.focus == 0 {
		model.focus = len(model.inputs) - 1
	} else {
		model.focus--
	}

	var cmd tea.Cmd
	model, cmd = model.changeFocusUpdate()

	return model, cmd
}

func (model *AddPool) changeFocusUpdate() (*AddPool, tea.Cmd) {
	cmds := make([]tea.Cmd, len(model.inputs))
	for i := range len(model.inputs) {
		if i == model.focus {
			// Set focused state
			cmds[i] = model.inputs[i].Focus()
			model.inputs[i].PromptStyle = focusedStyle
			model.inputs[i].TextStyle = focusedStyle
			continue
		}
		// Remove focused state
		model.inputs[i].Blur()
		model.inputs[i].PromptStyle = noStyle
		model.inputs[i].TextStyle = noStyle
	}

	return model, tea.Batch(cmds...)
}

func (model *AddPool) handleEnter() (*AddPool, tea.Cmd) {
	// if not last input selected
	if model.focus != len(model.inputs)-1 {
		// do nothing
		return model, nil
	}

	address := model.inputs[0].Value()
	network := model.inputs[1].Value()
	if address == "" || network == "" {
		model.warning = "It seems that one of the inputs is empty!"
		return model, clearWarning()
	}

	cmd := common.CmdHandler(AddPoolModeSubmitMsg{Address: address, Network: network})
	return model, cmd
}
func clearWarning() tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Second * 3)
		return clearWarningFieldMsg{}
	}
}

func (model *AddPool) handleSetSelectionMode() (tea.Model, tea.Cmd) {
	return model, common.CmdHandler(SelectionModeMsg{})
}

func NewAddPool() *AddPool {
	model := &AddPool{
		inputs: make([]textinput.Model, 2),
		keys:   keysAddPool,
	}

	var t textinput.Model
	for i := range model.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.Width = 30

		switch i {
		case 0:
			t.Placeholder = "Address (just CTRL + V)"
			t.Focus()
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.PlaceholderStyle = blurredStyle
		case 1:
			t.Placeholder = "Network (just CTRL + V)"
			t.PlaceholderStyle = blurredStyle
			t.CharLimit = 32
		}

		model.inputs[i] = t
	}

	return model
}
