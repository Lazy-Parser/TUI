package page_generator

import tea "github.com/charmbracelet/bubbletea"

type footer struct {
	stepList *StepList
}

func NewFooter() tea.Model {
	return &footer{
		stepList: NewStepList(),
	}
}

// add step
func addStep(title string) tea.Cmd {
	return func() tea.Msg {
		return AddStepMsg{Title: title}
	}
}

func (f *footer) Init() tea.Cmd {
	return tea.Batch(f.stepList.Init(), addStep("Mexc"))
}

func (f *footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	f.stepList, cmd = f.stepList.Update(msg)

	return f, cmd
}

func (f *footer) View() string {
	return f.stepList.View()
}
