package page_viewer

import tea "github.com/charmbracelet/bubbletea"

type footer struct{}

func NewFooter() tea.Model                                  { return &footer{} }
func (f *footer) Init() tea.Cmd                             { return nil }
func (f *footer) Update(msg tea.Msg) (tea.Model, tea.Cmd)   { return f, nil }
func (f *footer) View() string                              { return "Footer" }
