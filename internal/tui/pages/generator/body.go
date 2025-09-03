package page_generator

import (
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: make search by name / network
// !!!!!!!!!!!!!!!!!!!!!!!!!!
// TODO: create a page with the list of all running tasks
// Also make logic in some "service" and use only logic, not just core here
// !!!!!!!!!!!!!!!!!!!!!!!!!!

type Mode int

const (
	ModeSelection Mode = iota
	ModeAddPool
	ModeGeneration // maybe split generation to the "GenerationExchanges" and "GenerationManually"
)

type mainView struct {
	start      bool
	mode       Mode
	modeModels map[Mode]tea.Model
}

func NewMain() tea.Model {
	return &mainView{
		mode: ModeSelection,
		modeModels: map[Mode]tea.Model{
			ModeSelection:  NewSelection(),
			ModeAddPool:    NewAddPool(),
			ModeGeneration: NewGenetationMode(),
		},
	}
}

func (model *mainView) Init() tea.Cmd {
	// init selection mode, because it opens first. Init other modes, only when are open!

	return model.modeModels[model.mode].Init()
}

func (model *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case SelectionModeSubmitMsg:
		return model.handleSelectionSubmit(msg)
	case AddPoolModeSubmitMsg:
		return model.handleAddPoolSubmit(msg)
	case SelectionModeMsg:
		return model.handleSelectionMsg()
	case AddPoolModeMsg:
		return model.handleAddPoolMsg()

	case tea.KeyMsg:
		switch msg.String() {
		// exit
		case "q":
			return model, tea.Quit
		}
		switch msg.Type {
		case tea.KeyCtrlC:
			return model, tea.Quit
		}
	}

	// update selected mode
	var cmd tea.Cmd
	model.modeModels[model.mode], cmd = model.modeModels[model.mode].Update(msg)

	return model, cmd
}

func (model *mainView) View() string {
	return model.modeModels[model.mode].View()
}

// methods
func (m *mainView) handleSelectionSubmit(msg SelectionModeSubmitMsg) (tea.Model, tea.Cmd) {
	m.mode = ModeGeneration
	var str string
	for _, exchange := range msg.Elems {
		str += exchange + " | "
	}
	m.getGenerationMode().Address = ""
	m.getGenerationMode().Network = ""

	return m, nil
}
func (m *mainView) handleSelectionMsg() (tea.Model, tea.Cmd) {
	m.mode = ModeSelection
	return m, nil
}

func (m *mainView) handleAddPoolSubmit(msg AddPoolModeSubmitMsg) (tea.Model, tea.Cmd) {
	m.mode = ModeGeneration
	m.getGenerationMode().Set(msg.Address, msg.Network)

	// also do not forget to Init() this model to start generation process!
	cmd := m.getGenerationMode().Init()

	return m, cmd
}
func (m *mainView) handleAddPoolMsg() (tea.Model, tea.Cmd) {
	m.mode = ModeAddPool
	return m, nil
}

// helpers
func (m *mainView) getGenerationMode() *GeneratationMode {
	return m.modeModels[ModeGeneration].(*GeneratationMode)
}
