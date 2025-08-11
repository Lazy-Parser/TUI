package page_viewer

import (
	tea "github.com/charmbracelet/bubbletea"
)

type mainView struct{}

func NewMain() tea.Model                                   { return &mainView{} }
func (m *mainView) Init() tea.Cmd                          { return nil }
func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd){ return m, nil }
func (m *mainView) View() string                           { return "Content goes here" }
