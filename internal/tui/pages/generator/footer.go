package page_generator

import (
	"github.com/Lazy-Parser/TUI/internal/tui/common"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

type footer struct {
	mode Mode

	help          help.Model
	keySelection  keyMapSelection
	keyAddPool    keyMapAddPool
	keyGeneration keyMapGeneration
	// OR keys []
}

func NewFooter() tea.Model {
	return &footer{
		mode: ModeSelection,

		help: help.New(),
		keySelection:  keysSelection,
		keyAddPool:    keysAddPool,
		keyGeneration: keysGeneration,
	}
}

func (f *footer) Init() tea.Cmd {
	return nil
}

func (f *footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case SelectionModeSubmitMsg, AddPoolModeSubmitMsg:
		f.mode = ModeGenerationAddPool
		return f, nil

	case SelectionModeMsg:
		f.mode = ModeSelection
		return f, nil

	case AddPoolModeMsg:
		f.mode = ModeAddPool
		return f, nil
	}

	return f, nil
}

func (f *footer) View() string {
	var bindings []key.Binding
	switch f.mode {
	case ModeSelection:
		bindings = common.KeyMapToSlice(f.keySelection)
	case ModeAddPool:
		bindings = common.KeyMapToSlice(f.keyAddPool)
	case ModeGenerationAddPool:
		bindings = common.KeyMapToSlice(f.keyGeneration)
	}
	
	return f.help.ShortHelpView(bindings)
}
