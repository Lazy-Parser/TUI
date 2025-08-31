package page_generator

import tea "github.com/charmbracelet/bubbletea"

type GeneratationMode struct {
	Info string
}

func (model *GeneratationMode) Init() tea.Cmd {
	return nil
}

func (model *GeneratationMode) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return model, nil
}

func (model *GeneratationMode) View() string {
	return model.Info
}

func NewGenetationMode() *GeneratationMode {
	return &GeneratationMode{}
}
