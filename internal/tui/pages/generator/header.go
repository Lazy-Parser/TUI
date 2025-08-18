package page_generator

import (
	tea "github.com/charmbracelet/bubbletea"
)

type header struct{}

func NewHeader() tea.Model { return &header{} }

func (h *header) Init() tea.Cmd                            { return nil }
func (h *header) Update(msg tea.Msg) (tea.Model, tea.Cmd)  { return h, nil }
func (h *header) View() string                             { return "Header" }
