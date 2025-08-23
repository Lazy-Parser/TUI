package page_guide

import (
	tea "github.com/charmbracelet/bubbletea"
)

// TODO: write a guide about layout: 1 - header, 2 - footer, esc - to the main
// maybe rename "default" page to the "menu"
// write about generator, listener, logs, database (that we can view it)
// use https://github.com/charmbracelet/bubbles?tab=readme-ov-file#viewport for this

type mainView struct{}

func NewMain() tea.Model                                    { return &mainView{} }
func (m *mainView) Init() tea.Cmd                           { return nil }
func (m *mainView) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return m, nil }
func (m *mainView) View() string                            { return "Content goes here" }
